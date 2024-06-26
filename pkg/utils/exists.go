package utils

import (
	"errors"
	"os"
)

func IsExists(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}
