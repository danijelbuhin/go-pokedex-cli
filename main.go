package main

import (
	"time"

	"github.com/danijelbuhin/go-pokedex-cli/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaUrl *string
	previousLocationAreaUrl *string
	caughtPokemon map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}