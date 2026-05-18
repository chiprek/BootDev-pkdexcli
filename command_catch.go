package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {

	if len(args) != 1 {
		return fmt.Errorf("please provide a pokemon name")
	}

	resp, err := cfg.pokeapiClient.GetPokemon(args[0])
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", args[0])
	roll := rand.Intn(resp.BaseExperience)
	if roll <= 40 {
		cfg.caughtPokemon[resp.Name] = resp
		fmt.Printf("%s was caught!\n", resp.Name)
		fmt.Println("you may now inspect it with the inspect command")
	} else {
		fmt.Printf("%s has escaped!\n", resp.Name)
	}
	return nil
}
