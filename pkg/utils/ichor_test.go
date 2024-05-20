package utils

import (
	"testing"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/stretchr/testify/assert"
)

func TestIchorCategorize(t *testing.T) {
	artifacts := []api.Artifacts{
		{
			Name: "test classpath.jar",
			Type: "CLASS_PATH",
		},
		{
			Name: "test external.jar",
			Type: "EXTERNAL_FILE",
		},
		{
			Name: "test natives.jar",
			Type: "NATIVES",
		},
	}

	var meta api.LaunchMeta
	meta.LaunchTypeData.Artifacts = artifacts

	sorted := CategorizeArtifacts(meta, "test")

	assert.Equal(t, []string{"test/test classpath.jar"}, sorted["CLASS_PATH"])
	assert.Equal(t, []string{"test classpath.jar"}, sorted["ICHOR"])
	assert.Equal(t, []string{"test external.jar"}, sorted["EXTERNAL_FILE"])
	assert.Equal(t, []string{"test/test natives.jar"}, sorted["NATIVES"])
}
