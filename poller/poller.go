package poller

import (
	"io/fs"
	"os"
	"path/filepath"
	"time"
)

type PollerEvent struct {
	Path string
}

type PollerError error

type Poller struct {
	// Milliseconds
	interval      int
	directories   []string
	files         []string
	fileModsTimes map[string]int64
	Events        chan PollerEvent
	Errors        chan PollerError
	FS            fs.ReadDirFS
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

func (pw *Poller) Add(dir string) {
	filepath.WalkDir(dir, func(path string, d fs.DirEntry, err error) error {
		if d.IsDir() && dir != path {
			pw.directories = append(pw.directories, path)
			pw.Add(path)
		} else {
			pw.files = append(pw.files, path)
			fInfo, err := d.Info()
			if err != nil {
				pw.Errors <- err
			}
			pw.fileModsTimes[path] = fInfo.ModTime().UnixMilli()
		}
		return nil
	})
}

func (pw *Poller) StartPoller() chan struct{} {
	ticker := time.NewTicker(time.Duration(pw.interval) * time.Millisecond)
	quit := make(chan struct{})
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, filePath := range pw.files {
					oldFileTime := pw.fileModsTimes[filePath]
					fileInfo, err := os.Stat(filePath)
					if err != nil {
						continue
					}
					if fileInfo.ModTime().UnixMilli() != oldFileTime {
						pw.Events <- PollerEvent{
							Path: filePath,
						}
						pw.fileModsTimes[filePath] = fileInfo.ModTime().UnixMilli()
						break
					}
				}
			case <-quit:
				ticker.Stop()
				return
			}
		}
	}()

	return quit
}

func walkDir(fs fs.ReadDirFS, dir string, fn fs.WalkDirFunc) error {
	entries, err := fs.ReadDir(dir)
	if err != nil {
		return err
	}

	for _, entry := range entries {
		path := filepath.Join(dir, entry.Name())
		fn(path, entry, nil)
	}

	return nil
}
