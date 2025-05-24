package pokeapi

import (
	"encoding/json"
)

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`

	Stats []struct {
		Base int `json:"base_stat"`
		Stat struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"stats"`

	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

func (c *Client) GetPokemon(name string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + name
	dat, err := c.get(url)
	if err != nil {
		return Pokemon{}, err
	}
	var pkm Pokemon
	if err := json.Unmarshal(dat, &pkm); err != nil {
		return Pokemon{}, err
	}
	return pkm, nil
}
