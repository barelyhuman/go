package semver_test

import (
	"testing"

	"github.com/barelyhuman/go/semver"
)

func TestLatestVersion(t *testing.T) {
	t.Run("test ascending order", func(t *testing.T) {
		latest := semver.LatestVersion([]string{"v0.0.1", "v0.0.2", "v0.0.3"})
		if latest != "v0.0.3" {
			t.Fail()
		}
	})

	t.Run("test descending order", func(t *testing.T) {
		latest := semver.LatestVersion([]string{"v1.0.0", "v0.0.3", "v0.0.2", "v0.0.3"})
		if latest != "v1.0.0" {
			t.Fail()
		}
	})

	t.Run("test random order", func(t *testing.T) {
		latest := semver.LatestVersion([]string{"v0.0.0", "v0.0.2", "v1.0.0", "v0.0.1"})
		if latest != "v1.0.0" {
			t.Fail()
		}
	})
}
