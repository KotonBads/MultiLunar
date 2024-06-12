package api

import "testing"

func TestFetchMeta(t *testing.T) {
	launchbody := LaunchBody{
		OS:      "linux",
		Version: "1.8.9",
		Arch:    "x64",
		Module:  "forge",
	}

	res, err := launchbody.FetchLaunchMeta()
	if err != nil {
		t.Error(err)
	}

	if !res.Success {
		t.Error(res)
	}

	t.Log(res)
}
