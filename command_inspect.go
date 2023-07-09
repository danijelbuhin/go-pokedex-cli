package main

import (
	"errors"
	"fmt"
)

func callbackInspect(cfg *config, args ...string) error {
	pokemonName := args[0]
	
	if len(pokemonName) == 0 {
		return errors.New("please enter pokemon name")
	}

	pokemon, ok := cfg.caughtPokemon[pokemonName]

	if !ok {
		return fmt.Errorf("%s is not caught yet", pokemonName)
	}

	fmt.Printf("Here are some information about %s:\n", pokemonName)
	fmt.Printf("Base experience: %v \n", pokemon.BaseExperience)
	fmt.Printf("Height: %v \n", pokemon.Height)
	fmt.Printf("Species: %v \n", pokemon.Species.Name)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("--- %v: %v\n", stat.Stat.Name, stat.BaseStat)
	}

	fmt.Println("Types:")
	for _, typ := range pokemon.Types {
		fmt.Printf("--- %s\n", typ.Type.Name)
	}

	return nil
}