package download

import (
	"fmt"
	"path/filepath"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"golang.org/x/mod/semver"
)

/*
file structure:

	instances
	|-instance
	|	|-instance specific artifact(s)
	|
	|-common
		|-textures
		|-artifacts
		|-jre

updater:
	- update all artifacts
	- update all required artifacts
	- check natives

natives:
	- there are 2 kinds, presumably lwgl 2 and 3
	- possible just have this in common dir
	- change natives dir depending on mc ver
*/

func GetNativesDir(data api.LaunchBody, path string) (string, error) {
	if !semver.IsValid("v" + data.Version) {
		return "", fmt.Errorf("error: semver not valid: %v", data.Version)
	}

	out := filepath.Join(path, "modern")

	if semver.Compare("v"+data.Version, "v1.13") < 0 {
		out = filepath.Join(path, "legacy")
	}

	return out, nil
}
