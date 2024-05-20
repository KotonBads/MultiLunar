package utils

import (
	"fmt"

	"github.com/KotonBads/MultiLunar/pkg/api"
)

func CategorizeArtifacts(data api.LaunchMeta, path string) map[string][]string {
	artifacts := make(map[string][]string)

	for _, val := range data.LaunchTypeData.Artifacts {
		switch val.Type {
		case "CLASS_PATH":
			artifacts["CLASS_PATH"] = append(artifacts["CLASS_PATH"], fmt.Sprintf("%s/%s", path, val.Name))
			artifacts["ICHOR"] = append(artifacts["ICHOR"], val.Name)
		case "EXTERNAL_FILE":
			artifacts["EXTERNAL_FILE"] = append(artifacts["EXTERNAL"], val.Name)
		case "NATIVES":
			artifacts["NATIVES"] = append(artifacts["NATIVES"], fmt.Sprintf("%s/%s", path, val.Name))
		}
	}

	return artifacts
}
