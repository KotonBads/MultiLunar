package download

import (
	"fmt"
	"os"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/KotonBads/MultiLunar/pkg/utils"
	"github.com/c4milo/unpackit"
)

func UnpackNatives(data api.LaunchMeta, native string, path string) error {
	if !data.Success {
		return fmt.Errorf("error: data.Success is false")
	}

	out := utils.CategorizeArtifacts(data, native)["NATIVES"][0]

	file, err := os.Open(out)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()
	unpackit.Unpack(file, path)

	return nil
}
