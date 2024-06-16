package instance

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

/*

file structure:

	instances
	|-instance1
		|-instance.json
	|-instance2
	|-instance3

*/

func CreateInstance(name string, path string, version string) error {
	id := uuid.NewString()

	err := os.MkdirAll(filepath.Join(path, id), os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating dirs: %s", err)
	}

	file, err := os.Create(filepath.Join(path, id, "instance.json"))
	if err != nil {
		return fmt.Errorf("error creating file: %s", err)
	}
	defer file.Close()

	instance := &Instance{
		DisplayName: name,
		Path:        filepath.Join(path, id),
		Version:     version,
		Id:          id,
	}

	out, err := json.MarshalIndent(instance, "", "\t")
	if err != nil {
		return fmt.Errorf("error marshalling instance info: %s", err)
	}

	written, err := file.Write(out)
	if err != nil {
		return fmt.Errorf("error writing to file: %s", err)
	}

	log.Printf("[INFO] Written %d bytes to file", written)

	return nil
}
