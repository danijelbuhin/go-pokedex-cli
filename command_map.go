package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config) error {
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaUrl)

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")	
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous

	return nil
}

func callbackMapb(cfg *config) error {
	if cfg.previousLocationAreaUrl == nil {
		return errors.New("You are on the first page")
	}
	
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.previousLocationAreaUrl)

	if err != nil {
		return err
	}

	fmt.Println("Location areas:")	
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaUrl = resp.Next
	cfg.previousLocationAreaUrl = resp.Previous

	return nil
}