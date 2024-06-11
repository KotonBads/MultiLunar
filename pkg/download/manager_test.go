package download

import (
	"testing"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestNativeDir(t *testing.T) {
	body := api.LaunchBody{
		OS:      "linux",
		Version: "1.8.9",
		Module:  "forge",
		Arch:    "x64",
	}

	out, err := GetNativesDir(body, "testing")
	assert.Equal(t, nil, err)
	assert.Equal(t, "testing/legacy", out)

	body.Version = "1.16.1"

	out, err = GetNativesDir(body, "testing")
	assert.Equal(t, nil, err)
	assert.Equal(t, "testing/modern", out)
}

func TestDownloadAll(t *testing.T) {
	body := api.LaunchBody{
		OS:      "linux",
		Version: "1.8.9",
		Module:  "forge",
		Arch:    "x64",
	}

	meta, err := body.FetchLaunchMeta()
	assert.Equal(t, nil, err)

	failures, err := DownloadAll(meta, "../../temp/")
	assert.Equal(t, nil, err)
	assert.Equal(
		t,
		map[string]int{
			"artifacts": 0,
			"textures":  0,
		},
		failures,
	)
}
