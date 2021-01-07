package metadata

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const BaseURL = "https://metadata.platformequinix.com"

func GetMetadata() (*CurrentDevice, error) {
	return GetMetadataFromURL(BaseURL)
}

func GetMetadataFromURL(baseURL string) (*CurrentDevice, error) {
	res, err := http.Get(baseURL + "/metadata")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		return nil, err
	}

	var result struct {
		Error string `json:"error"`
		*CurrentDevice
	}
	if err := json.Unmarshal(b, &result); err != nil {
		if res.StatusCode >= 400 {
			return nil, errors.New(res.Status)
		}
		return nil, err
	}
	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return result.CurrentDevice, nil
}

func GetUserData() ([]byte, error) {
	return GetUserDataFromURL(BaseURL)
}

func GetUserDataFromURL(baseURL string) ([]byte, error) {
	res, err := http.Get(baseURL + "/userdata")
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	return b, err
}
