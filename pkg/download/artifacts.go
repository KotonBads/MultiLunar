package download

import (
	"fmt"
	"log"
	"path/filepath"
	"sync"
	"time"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/KotonBads/MultiLunar/pkg/utils"
)

func DownloadArtifacts(data api.LaunchMeta, path string) (failures int, err error) {
	if !data.Success {
		return 0, fmt.Errorf("error: data.Success is false")
	}

	var wg sync.WaitGroup
	var failed []int
	failures = 0

	for idx, val := range data.LaunchTypeData.Artifacts {
		wg.Add(1)

		go func(idx int, artifact api.Artifacts) {
			defer wg.Done()

			out := filepath.Join(path, artifact.Name)

			if utils.IsExists(out) && utils.HashIsEqual(out, artifact.Sha1) {
				log.Printf("[INFO] Already up-to-date : %s\n", artifact.Name)
				return
			}

			err = utils.DownloadFile(out, artifact.Url, time.Second*30)
			if err != nil {
				log.Printf("[WARN] Error downloading [%s] : %s\n", artifact.Name, err)
				failed = append(failed, idx)
				failures++
				return
			}

			log.Printf("[INFO] Downloaded artifact : %s\n", artifact.Name)
		}(idx, val)
	}
	wg.Wait()

	if failures > 0 {
		log.Printf("[INFO] %d failures, attempting to download", failures)

		for _, val := range failed {
			artifact := data.LaunchTypeData.Artifacts[val]

			out := filepath.Join(path, artifact.Name)

			if utils.IsExists(out) && utils.HashIsEqual(out, artifact.Sha1) {
				log.Printf("[INFO] Already up-to-date : %s\n", artifact.Name)
				failures--
				return
			}

			err = utils.DownloadFile(path, artifact.Url, time.Second*30)
			if err != nil {
				log.Printf("[WARN] Error downloading artifact : %s\n", artifact.Name)
				return
			}

			log.Printf("[INFO] Downloaded artifact : %s\n", artifact.Name)
			failures--
		}
	}

	log.Printf("[INFO] Downloaded artifacts with %d failures", failures)

	return failures, nil
}
