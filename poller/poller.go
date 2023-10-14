package poller

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type PollerEvent struct {
	Path string
}

type PollerError error

type EventHandler func(e PollerEvent)
type ErrorHandler func(e PollerError)

type Poller struct {
	// Milliseconds
	interval      int
	directories   []string
	files         []string
	fileModsTimes map[string]int64
	filesToIgnore []string
	Events        chan PollerEvent
	evntHandlers  []EventHandler
	errHandlers   []ErrorHandler
	Errors        chan PollerError
	quit          chan struct{}
}

func NewPollWatcher(
	interval int,
) *Poller {
	poller := &Poller{
		Events:        make(chan PollerEvent, 1),
		fileModsTimes: map[string]int64{},
		interval:      interval,
	}
	return poller
}

func (pw *Poller) IgnorePattern(pattern string) error {
	matches, err := filepath.Glob(pattern)
	fmt.Printf("matches: %v\n", matches)
	if err != nil {
		return err
	}

	pw.filesToIgnore = append(pw.filesToIgnore, matches...)
	return nil
}

func (pw *Poller) IgnorePath(path string) {
	pw.filesToIgnore = append(pw.filesToIgnore, filepath.Clean(path))
}

func (pw *Poller) isIgnored(path string) bool {
	for _, p := range pw.filesToIgnore {
		if strings.HasPrefix(path, p) || path == p {
			return true
		}
	}
	return false
}

func (pw *Poller) AddPath(path string) {
	cleanPath := filepath.Clean(path)
	if pw.isIgnored(cleanPath) {
		return
	}

	statInfo, err := os.Stat(path)
	if err != nil {
		pw.Errors <- err
		return
	}

	pw.files = append(pw.files, cleanPath)
	pw.fileModsTimes[cleanPath] = statInfo.ModTime().UnixMilli()
}

func (pw *Poller) Add(dir string) {
	if pw.isIgnored(dir) {
		return
	}
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && dir != path {
			pw.directories = append(pw.directories, path)
			pw.Add(path)
		} else {
			pw.AddPath(path)
		}
		return nil
	})
}

func (pw *Poller) OnError(fn func(e PollerError)) {
	go func() {
		for err := range pw.Errors {
			fn(err)
		}
	}()
}

func (pw *Poller) OnEvent(fn func(e PollerEvent)) {
	pw.evntHandlers = append(pw.evntHandlers, fn)
}

func (pw *Poller) Stop() {
	pw.quit <- struct{}{}
}

func (pw *Poller) Wait() chan struct{} {
	return pw.quit
}

func (pw *Poller) Start() {
	ticker := time.NewTicker(time.Duration(pw.interval) * time.Millisecond)
	pw.quit = make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, filePath := range pw.files {
					if pw.isIgnored(filePath) {
						continue
					}
					oldFileTime := pw.fileModsTimes[filePath]
					fileInfo, err := os.Stat(filePath)
					if err != nil {
						continue
					}
					if fileInfo.ModTime().UnixMilli() != oldFileTime {
						pollerEvnt := PollerEvent{
							Path: filePath,
						}
						for _, eh := range pw.evntHandlers {
							eh(pollerEvnt)
						}
						pw.Events <- pollerEvnt
						pw.fileModsTimes[filePath] = fileInfo.ModTime().UnixMilli()
						break
					}
				}
			case <-pw.quit:
				ticker.Stop()
				return
			}
		}
	}()

}
