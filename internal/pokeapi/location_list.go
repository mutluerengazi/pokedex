package pokeapi

import (
	"encoding/json"
)

// ListLocations -
func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	dat, err := c.get(url)
	if err != nil {
		return RespShallowLocations{}, err
	}

	var out RespShallowLocations
	if err := json.Unmarshal(dat, &out); err != nil {
		return RespShallowLocations{}, err
	}
	return out, nil
}
