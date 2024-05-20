package download

import (
	"testing"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestDownloadTextures(t *testing.T) {
	body := api.LaunchBody{
		OS: "linux",
		Version: "1.8.9",
		Module: "forge",
		Arch: "x64",
	}

	meta, err := body.FetchLaunchMeta()
	assert.Equal(t, nil, err)

	status, err := DownloadTextures(meta, "/home/koton-bads/Documents/Go/MultiLunar/temp/")
	assert.Equal(t, nil, err)
	assert.Equal(t, 0, status)
}