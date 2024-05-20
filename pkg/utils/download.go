package utils

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func DownloadFile(path string, url string, timeout time.Duration) error {
	err := os.MkdirAll(filepath.Base(path), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating dirs: %s", err)
	}

	out, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("error creating [%s]: %s", path, err)
	}
	defer out.Close()

	client := http.Client{
		Timeout: timeout,
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("error sending request: %s", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error copying bytes: %s", err)
	}

	return nil
}
