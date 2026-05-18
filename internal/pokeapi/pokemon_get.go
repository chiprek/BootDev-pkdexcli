package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	pokemonURL := baseURL + "/pokemon/" + pokemonName

	if cached, ok := c.cache.Get(pokemonURL); ok == true {
		data := Pokemon{}
		json.Unmarshal(cached, &data)
		return data, nil
	}

	req, err := http.NewRequest("GET", pokemonURL, nil)
	if err != nil {
		return Pokemon{}, fmt.Errorf("GET request failed")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Pokemon{}, fmt.Errorf("Failed to read response")
	}

	data := Pokemon{}
	if err := json.Unmarshal(body, &data); err != nil {
		return Pokemon{}, err
	}
	c.cache.Add(pokemonURL, body)
	return data, nil
}
