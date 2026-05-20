package main

import (
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("please provide a Pokemon's name\n")
	}
	pokemon, ok := cfg.caughtPokemon[args[0]]
	if ok == false {
		return fmt.Errorf("you have not caught that pokemon\n")
	} else {
		fmt.Printf("Name: %s\n", pokemon.Name)
		fmt.Printf("Height: %d\n", pokemon.Height)
		fmt.Printf("Weight: %d\n", pokemon.Weight)
		fmt.Println("Stats:")

		for _, stat := range pokemon.Stats {
			fmt.Printf(" - %s: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, typeInfo := range pokemon.Types {
			fmt.Printf(" - %s\n", typeInfo.Type.Name)
		}
	}
	return nil
}
