package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (Pokemon, error) {
	endpoint := "/pokemon/" + pokemon
	fullURL := baseUrl + endpoint

	data, ok := c.cache.Get(fullURL)

	if ok {
		fmt.Println("Getting from cache")
		poke := Pokemon{}

		err := json.Unmarshal(data, &poke)
	
		if err != nil {
			return Pokemon{}, err
		}
	
		return poke, nil
	}
	
	req, err := http.NewRequest("GET", fullURL, nil)

	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)

	if err != nil {
		return Pokemon{}, err
	}

	defer resp.Body.Close()

	if resp.StatusCode > 399 {
		return Pokemon{}, fmt.Errorf("bad status code: %v", resp.StatusCode)
	}

	data, err = io.ReadAll(resp.Body)

	if err != nil {
		return Pokemon{}, err
	}

	poke := Pokemon{}

	err = json.Unmarshal(data, &poke)

	if err != nil {
		return Pokemon{}, err
	}

	c.cache.Add(fullURL, data)

	return poke, nil
}