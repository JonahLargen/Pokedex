package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaInfoResponse struct {
	ID                   int                   `json:"id"`
	Name                 string                `json:"name"`
	GameIndex            int                   `json:"game_index"`
	EncounterMethodRates []EncounterMethodRate `json:"encounter_method_rates"`
	Location             NamedAPIResource      `json:"location"`
	Names                []Name                `json:"names"`
	PokemonEncounters    []PokemonEncounter    `json:"pokemon_encounters"`
}

type EncounterMethodRate struct {
	EncounterMethod NamedAPIResource               `json:"encounter_method"`
	VersionDetails  []EncounterMethodVersionDetail `json:"version_details"`
}

type EncounterMethodVersionDetail struct {
	Rate    int              `json:"rate"`
	Version NamedAPIResource `json:"version"`
}

type NamedAPIResource struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type Name struct {
	Name     string           `json:"name"`
	Language NamedAPIResource `json:"language"`
}

type PokemonEncounter struct {
	Pokemon        NamedAPIResource       `json:"pokemon"`
	VersionDetails []PokemonVersionDetail `json:"version_details"`
}

type PokemonVersionDetail struct {
	MaxChance        int               `json:"max_chance"`
	Version          NamedAPIResource  `json:"version"`
	EncounterDetails []EncounterDetail `json:"encounter_details"`
}

type EncounterDetail struct {
	Chance          int                `json:"chance"`
	ConditionValues []NamedAPIResource `json:"condition_values"`
	MaxLevel        int                `json:"max_level"`
	Method          NamedAPIResource   `json:"method"`
	MinLevel        int                `json:"min_level"`
}

func GetLocationAreaInfo(locationArea string) (*LocationAreaInfoResponse, error) {
	url := baseURL + "/location-area/" + locationArea
	if cachedData, found := pokeCache.Get(url); found {
		var cachedResp LocationAreaInfoResponse
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
	if resp.StatusCode == http.StatusNotFound {
		return nil, fmt.Errorf("location area '%s' not found", locationArea)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch data: status %s", resp.Status)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	pokeCache.Add(url, body)
	var apiResp LocationAreaInfoResponse
	if err := json.Unmarshal(body, &apiResp); err != nil {
		return nil, err
	}
	return &apiResp, nil
}
