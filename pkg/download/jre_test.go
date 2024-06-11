package download

import (
	"testing"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestJRE(t *testing.T) {
	body := api.LaunchBody{
		OS:      "linux",
		Version: "1.8.9",
		Module:  "forge",
		Arch:    "x64",
	}

	meta, err := body.FetchLaunchMeta()
	assert.Equal(t, nil, err)

	err = DownloadJRE(meta, "../../temp/jre")
	assert.Equal(t, nil, err)
}
