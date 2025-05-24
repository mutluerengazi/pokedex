package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func catchProbability(baseExp int) int {
	chance := 100 - baseExp
	if chance < 10 {
		chance = 10
	}
	return chance
}

func commandCatch(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>")
	}
	name := args[0]

	if _, ok := cfg.pokedex[name]; ok {
		fmt.Printf("%s is already in your Pokedex!\n", name)
		return nil
	}

	pkm, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	prob := catchProbability(pkm.BaseExperience)
	roll := rand.Intn(100) // default RNG is already seeded in Go 1.20+

	if roll < prob {
		cfg.pokedex[name] = pkm
		fmt.Printf("%s was caught!\n", name)
	} else {
		fmt.Printf("%s escaped!\n", name)
	}
	return nil
}
