package poller

import (
	"os"
	"path/filepath"
	"testing"
)

func TestAddition(t *testing.T) {
	t.Run("test basic path addition", func(t *testing.T) {
		poller := NewPollWatcher(2000)
		poller.AddPath("./poller.go")
		if len(poller.files) != 1 {
			t.Fail()
		}

		if poller.files[0] != "poller.go" {
			t.Fail()
		}
	})

	t.Run("test basic directory addition", func(t *testing.T) {
		poller := NewPollWatcher(2000)
		poller.Add(".")
		if len(poller.files) == 0 {
			t.Fail()
		}
	})

	t.Run("test ignored addition", func(t *testing.T) {
		poller := NewPollWatcher(2000)
		poller.IgnorePath("./poller.go")
		poller.Add(".")

		if !poller.isIgnored("./poller.go") {
			t.Fail()
		}

		if !poller.isIgnored("poller.go") {
			t.Fail()
		}

		for _, v := range poller.files {
			if v == filepath.Clean("./poller.go") {
				t.Fail()
			}
		}
	})
}

func TestEvents(t *testing.T) {
	t.Run("basic events", func(t *testing.T) {
		poller := NewPollWatcher(100)
		poller.Add(".")

		poller.OnEvent(func(e PollerEvent) {
			if e.Path != "poller.go" {
				t.Fail()
			}
			poller.Stop()
		})

		poller.Start()

		go func() {
			file, err := os.ReadFile("./poller.go")
			if err != nil {
				t.Logf(err.Error())
				t.Fail()
			}
			err = os.WriteFile("./poller.go", file, os.ModePerm)
			if err != nil {
				t.Logf(err.Error())
				t.Fail()
			}
		}()

		<-poller.Wait()
	})
}
