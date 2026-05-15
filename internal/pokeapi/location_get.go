package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetLocation(locationName string) (Location, error) {
	locationURL := baseURL + "/location-area/" + locationName

	if cached, ok := c.cache.Get(locationURL); ok == true {
		data := Location{}
		json.Unmarshal(cached, &data)
		return data, nil
	}

	req, err := http.NewRequest("GET", locationURL, nil)
	if err != nil {
		return Location{}, fmt.Errorf("Get request unsuccessful")
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, fmt.Errorf("Failed to read response")
	}

	data := Location{}

	if err := json.Unmarshal(body, &data); err != nil {
		return Location{}, err
	}

	c.cache.Add(locationURL, body)
	return data, nil
}
