package utils

import (
	"crypto/sha1"
	"io"
	"os"
)

func HashIsEqual(path string, hash string) (eq bool) {
	file, err := os.Open(path)
	if err != nil {
		return false
	}
	defer file.Close()

	h := sha1.New()

	_, err = io.Copy(h, file)
	if err != nil {
		return false
	}

	if string(h.Sum(nil)) != hash {
		return false
	}

	return true
}
