package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int                  `json:"count"`
	Next     string               `json:"next"`
	Previous string               `json:"previous"`
	Results  []LocationAreaResult `json:"results"`
}

type LocationAreaResult struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

func GetLocationAreas(pageUrl string) (*LocationAreaResponse, error) {
	url := baseURL + "/location-area"
	if pageUrl != "" {
		url = pageUrl
	}
	if cachedData, found := pokeCache.Get(url); found {
		var cachedResp LocationAreaResponse
		if err := json.Unmarshal(cachedData, &cachedResp); err != nil {
			return nil, fmt.Errorf("failed to unmarshal cached data: %w", err)
		}
		return &cachedResp, nil
	}
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pokeCache.Add(url, body)
	var apiResp LocationAreaResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}
	return &apiResp, nil
}
