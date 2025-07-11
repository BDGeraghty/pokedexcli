package main

import (
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet.")
		return nil
	}

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("  - %s\n", pokemon.Name)
	}

	return nil
}
