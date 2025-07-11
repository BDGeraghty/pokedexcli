package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name")
	}

	name := args[0]
	pokemon, exists := cfg.caughtPokemon[name]
	if !exists {
		return fmt.Errorf("you have not caught a %s", name)
	}

	fmt.Printf("Inspecting %s...\n", pokemon.Name)
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, typeInfo := range pokemon.Types {
		fmt.Printf("  - %s\n", typeInfo.Type.Name)
	}

	return nil
}
/*
Name: pidgey
Height: 3
Weight: 18
Stats:
  -hp: 40
  -attack: 45
  -defense: 40
  -special-attack: 35
  -special-defense: 35
  -speed: 56
Types:
  - normal
  - flying
  */