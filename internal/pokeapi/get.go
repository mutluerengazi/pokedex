package pokeapi

import (
	"io"
	"net/http"
)

func (c *Client) get(url string) ([]byte, error) {
	if val, ok := c.cache.Get(url); ok {
		return val, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	c.cache.Add(url, dat)
	return dat, nil
}
