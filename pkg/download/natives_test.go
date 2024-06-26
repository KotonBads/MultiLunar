package download

import (
	"testing"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestUnpackNatives(t *testing.T) {
	body := api.LaunchBody{
		OS:      "linux",
		Version: "1.8.9",
		Module:  "forge",
		Arch:    "x64",
	}

	meta, err := body.FetchLaunchMeta()
	assert.Equal(t, nil, err)

	status, err := DownloadArtifacts(meta, "../../temp/artifacts")
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, status)

	out, err := GetNativesDir(body, "../../temp/natives")
	assert.Equal(t, nil, err)
	assert.Equal(t, "../../temp/natives/legacy", out)

	err = UnpackNatives(meta, "../../temp/artifacts", out)
	assert.Equal(t, nil, err)
}
