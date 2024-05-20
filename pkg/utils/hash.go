package utils

import (
	"crypto/sha1"
	"fmt"
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

	return fmt.Sprintf("%x", h.Sum(nil)) == hash
}
