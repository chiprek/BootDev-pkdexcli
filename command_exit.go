package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	fmt.Println("=======================================")
	os.Exit(0)
	return nil
}
