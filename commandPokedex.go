package main

import "fmt"

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return fmt.Errorf("did you mean to inspect pokedex? type only pokedex\n")
	}
	if len(cfg.caughtPokemon) == 0 {
		return fmt.Errorf("No caught pokemon")
	}
	fmt.Println("Your Pokedex:")
	for name := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}
	return nil
}
