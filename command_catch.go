package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func callbackCatch(cfg *config, args ...string) error {
	if len(args) == 0 {
		return errors.New("no pokemon provided")
	}

	pokemonName := args[0]
	
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)

	if err != nil {
		return err
	}
	const threshold = 50
	randNum := rand.Intn(pokemon.BaseExperience)
	fmt.Println(pokemon.BaseExperience, randNum, threshold)

	if randNum > threshold {
		return fmt.Errorf("failed to catch %s", pokemon.Name)
	}

	cfg.caughtPokemon[pokemonName] = pokemon
	fmt.Printf("%s is caught!\n", pokemon.Name)	

	return nil
}