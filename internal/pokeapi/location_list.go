package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	if c.cache != nil {
		if cached, ok := c.cache.Get(url); ok {
			var locations RespShallowLocations

			if err := json.Unmarshal(cached, &locations); err != nil {
				return RespShallowLocations{}, err
			}
			return locations, nil
		}
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespShallowLocations{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return RespShallowLocations{}, err
	}

	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return RespShallowLocations{}, err
	}

	locationsResponse := RespShallowLocations{}
	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return RespShallowLocations{}, err
	}

	return locationsResponse, nil
}
