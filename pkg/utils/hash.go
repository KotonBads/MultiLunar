package utils

import (
	"crypto/sha1"
	"fmt"
	"io"
	"os"
)

func HashIsEqual(path string, hash string) (eq bool, err error) {
	file, err := os.Open(path)
	if err != nil {
		return false, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	h := sha1.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return false, fmt.Errorf("error copying bytes: %s", err)
	}

	if string(h.Sum(nil)) != hash {
		return false, nil
	}

	return true, nil
}
