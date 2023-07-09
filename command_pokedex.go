package main

import (
	"errors"
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you haven't caught any pokemons yet")
	}

	fmt.Println("Your pokedex:")

	for _, poke := range cfg.caughtPokemon {
		fmt.Printf("-> %s \n", poke.Name)
	}
	
	return nil
}