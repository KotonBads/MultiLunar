package download

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/KotonBads/MultiLunar/pkg/utils"
	"github.com/c4milo/unpackit"
)

func DownloadJRE(data api.LaunchMeta, path string) error {
	if !data.Success {
		return fmt.Errorf("error: data.Success is false")
	}

	out := filepath.Join(path, data.Jre.ExecutablePathInArchive[0]+"."+data.Jre.Download.Extension)

	err := utils.DownloadFile(out, data.Jre.Download.URL, time.Second*120)
	if err != nil {
		return fmt.Errorf("error downloading jre: %s", err)
	}
	log.Printf("[INFO] Downloaded JRE")

	file, err := os.Open(out)
	if err != nil {
		return fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()
	unpackit.Unpack(file, path)

	return nil
}
