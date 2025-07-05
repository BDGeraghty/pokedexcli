package main

import (
	"time"

	"github.com/bootdotdev/pokedexcli/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(
		time.Second*10, // HTTP timeout
		time.Minute*5,   // Cache TTL
	)
	cfg := &config{
		pokeapiClient: pokeClient,
	}

	startRepl(cfg)
}
