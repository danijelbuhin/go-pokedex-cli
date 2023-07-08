package main

import "github.com/danijelbuhin/go-pokedex-cli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocationAreaUrl *string
	previousLocationAreaUrl *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}