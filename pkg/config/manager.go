package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func (config *LauncherConfig) LoadConfig(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("error opening config: %s", err)
	}
	defer file.Close()

	body, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error reading config: %s", err)
	}

	err = json.Unmarshal(body, config)
	if err != nil {
		return fmt.Errorf("error unmarshalling config: %s", err)
	}

	return nil
}

func (config *LauncherConfig) SaveConfig(path string) (int, error) {
	file, err := os.Create(path)
	if err != nil {
		return 0, fmt.Errorf("error opening file: %s", err)
	}
	defer file.Close()

	new, err := json.MarshalIndent(config, "", "\t")
	if err != nil {
		return 0, fmt.Errorf("error marshalling new config: %s", err)
	}

	written, err := file.Write(new)
	if err != nil {
		return 0, fmt.Errorf("error writing to file: %s", err)
	}

	return written, nil
}
