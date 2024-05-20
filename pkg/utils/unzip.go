package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func Unzip(src string, dest string) (err error) {
	reader, err := zip.OpenReader(src)
	if err != nil {
		return fmt.Errorf("error reading zip: %s", err)
	}
	defer reader.Close()

	err = os.MkdirAll(dest, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating dirs: %s", err)
	}

	for _, file := range reader.File {
		contents, err := file.Open()
		if err != nil {
			return fmt.Errorf("error reading file: %s", err)
		}
		defer contents.Close()

		path := filepath.Join(dest, file.Name)

		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return fmt.Errorf("error creating dirs: %s", err)
			}		
			break
		} 
	
		out, err := os.Create(path)
		if err != nil {
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, contents)
		if err != nil {
			return fmt.Errorf("error copying bytes: %s", err)
		}
	}

	return nil
}
