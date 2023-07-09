package main

import (
	"errors"
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
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

func callbackMapb(cfg *config, args ...string) error {
	if cfg.previousLocationAreaUrl == nil {
		return errors.New("you are on the first page")
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