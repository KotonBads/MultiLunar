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
