package main

import (
	"errors"
	"fmt"
)

func callbackExplore(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no location area provided")
	}

	locationAreaName := args[0]
	
	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)

	if err != nil {
		return err
	}

	fmt.Printf("Pokemon in %s: \n", locationArea.Name)	

	if len(locationArea.PokemonEncounters) == 0 {
		fmt.Println("No pokemon encouters in this location area")

		return nil
	}

	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}

	return nil
}