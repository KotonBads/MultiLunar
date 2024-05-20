package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"

	"github.com/google/uuid"
)

func (launchdata LaunchBody) FetchLaunchMeta() (res LaunchMeta, err error) {
	const url = "https://api.lunarclientprod.com/launcher/launch"
	var launchmeta LaunchMeta

	params := map[string]string{
		"hwid":             "0",
		"launch_type":      "OFFLINE",
		"branch":           "master",
		"launcher_version": "3.0.0",
		"os_release":       fmt.Sprintf("%d.%d.%d", rand.Intn(10), rand.Intn(10), rand.Intn(10)),
		"installation_id":  uuid.New().String(),
		"os":               launchdata.OS,
		"arch":             launchdata.Arch,
		"version":          launchdata.Version,
		"module":           launchdata.Module,
	}

	jsonParams, err := json.Marshal(params)
	if err != nil {
		return launchmeta, fmt.Errorf("error marshalling params: %s", err)
	}

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonParams))
	if err != nil {
		return launchmeta, fmt.Errorf("error creating request: %s", err)
	}
	request.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return launchmeta, fmt.Errorf("error sending request: %s", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return launchmeta, fmt.Errorf("error reading body: %s", err)
	}

	err = json.Unmarshal(body, &launchmeta)
	if err != nil {
		return launchmeta, fmt.Errorf("error unmarshalling: %s", err)
	}

	return launchmeta, nil
}
