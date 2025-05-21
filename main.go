package main

import (
	"time"

	"github.com/mutluerengazi/pokedex/internal/pokeapi"
)

func main() {
	cfg := &config{
		pokeapiClient: pokeapi.NewClient(
			30*time.Second, // HTTP timeout
			5*time.Second,  // cache expiry
		),
	}

	startRepl(cfg)
}
