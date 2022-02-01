package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func loadJSONData(
	httpClient *http.Client,
	url string,
	responseData interface{},
) error {
	request, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return fmt.Errorf("unable to create the request: %w", err)
	}

	response, err := httpClient.Do(request)
	if err != nil {
		return fmt.Errorf("unable to send the request: %w", err)
	}
	defer response.Body.Close()

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return fmt.Errorf("unable to read the response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("request failed: %d %s", response.StatusCode, responseBytes)
	}

	if err := json.Unmarshal(responseBytes, responseData); err != nil {
		return fmt.Errorf("unable to unmarshal the response body: %w", err)
	}

	return nil
}
