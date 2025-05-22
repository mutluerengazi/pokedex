package pokeapi

import (
	"encoding/json"
)

type LocationArea struct {
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func (c *Client) GetLocationArea(areaName string) (LocationArea, error) {
	url := baseURL + "/location-area/" + areaName
	dat, err := c.get(url)
	if err != nil {
		return LocationArea{}, err
	}

	var area LocationArea
	if err := json.Unmarshal(dat, &area); err != nil {
		return LocationArea{}, err
	}
	return area, nil
}
