package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	//fmt.Printf("DEBUG: args length is %d, args are: '%q'\n", len(args), args)

	if len(args) != 1 {
		return fmt.Errorf("please provide a location name")
	}
	fmt.Printf("Exploring %v...\n", args[0])
	resp, err := cfg.pokeapiClient.GetLocation(args[0])
	if err != nil {
		return fmt.Errorf("unable to get location")
	}

	fmt.Println("Found Pokemon:")
	for _, encounters := range resp.PokemonEncounters {
		fmt.Printf(" - %s\n", encounters.Pokemon.Name)
	}
	return nil
}
