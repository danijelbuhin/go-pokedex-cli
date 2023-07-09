package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreasResponse, error) {
	endpoint := "/location"
	fullURL := baseUrl + endpoint

	if pageUrl != nil {
		fullURL = *pageUrl
 	}

	data, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("Getting from cache")
		locationAreas := LocationAreasResponse{}

		err := json.Unmarshal(data, &locationAreas)
	
		if err != nil {
			return LocationAreasResponse{}, err
		}
	
		return locationAreas, nil
	}
	
	
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return LocationAreasResponse{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreas := LocationAreasResponse{}

	err = json.Unmarshal(data, &locationAreas)

	if err != nil {
		return LocationAreasResponse{}, err
	}

	c.cache.Add(fullURL, data)

	return locationAreas, nil
}