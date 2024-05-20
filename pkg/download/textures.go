package download

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/KotonBads/MultiLunar/pkg/api"
	"github.com/KotonBads/MultiLunar/pkg/utils"
)

func DownloadTextures(data api.LaunchMeta, path string) (failures int, err error) {
	if !data.Success {
		return 0, fmt.Errorf("error: data.Success is false")
	}

	var wg sync.WaitGroup
	var failed []int
	failures = 0

	response, err := http.Get(data.Textures.IndexURL)
	if err != nil {
		return 0, fmt.Errorf("error sending request: %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return 0, fmt.Errorf("error reading body: %s", err)
	}

	index := strings.Split(string(body), "\n")

	for idx, val := range index {
		wg.Add(1)
		go func (idx int, texture string)  {
			defer wg.Done()

			// cols[0] -> file path
			// cols[1] -> file hash
			cols := strings.Split(texture, " ")
			out := filepath.Join(path, cols[0])

			if utils.IsExists(out) && utils.HashIsEqual(out, cols[1]) {
				log.Printf("[INFO] Already up-to-date : %s", cols[0])
				return
			}

			err = utils.DownloadFile(out, data.Textures.BaseURL+cols[1], time.Second*30)
			if err != nil {
				log.Printf("[WARN] Error downloading [%s] : %s", cols[0], err)
				failed = append(failed, idx)
				failures++
				return
			}

			log.Printf("[INFO] Downloaded texture : %s", cols[0])
		}(idx, val)
	}
	wg.Wait()

	if failures > 0 {
		for _, val := range failed {
			texture := index[val]
			cols := strings.Split(texture, " ")
			out := filepath.Join(path, cols[0])

			if utils.IsExists(out) && utils.HashIsEqual(out, cols[1]) {
				log.Printf("[INFO] Already up-to-date : %s", cols[0])
				failures--
				return
			}

			err = utils.DownloadFile(out, data.Textures.BaseURL+cols[1], time.Second*30)
			if err != nil {
				log.Printf("[WARN] Error downloading [%s] : %s", cols[0], err)
				return
			}

			log.Printf("[INFO] Downloaded texture : %s", cols[0])
			failures--
		}
	}

	log.Printf("[INFO] Downloaded textures with %d failures", failures)

	return failures, nil
}
