package config

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	expect := &LauncherConfig{
		InstancePath: "/path/to/instance",
		AssetPath:    "/path/to/assets",
		CommonPath:   "/path/to/common",
		JREPath:      "/path/to/jre",
		JREMax:       3072,
		JREMin:       1024,
	}

	config := &LauncherConfig{}

	err := config.LoadConfig("../../temp/config.json")
	assert.Equal(t, nil, err)
	assert.Equal(t, expect, config)
}

func TestSaveConfig(t *testing.T) {
	config := &LauncherConfig{
		InstancePath: "/path/to/instance",
		AssetPath:    "/path/to/assets",
		CommonPath:   "/path/to/common",
		JREPath:      "/path/to/jre",
		JREMax:       3072,
		JREMin:       1024,
	}

	written, err := config.SaveConfig("../../temp/config.json")
	assert.Equal(t, nil, err)
	assert.Greater(t, written, 0)
}
