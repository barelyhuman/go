package semver

import "golang.org/x/mod/semver"

func LatestVersion(versions []string) (latest string) {
	for _, v := range versions {
		if semver.IsValid(v) && semver.Compare(v, latest) == 1 {
			latest = v
		}
	}
	return
}
