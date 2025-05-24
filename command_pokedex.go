package main

import (
	"fmt"
	"sort"
)

func commandPokedex(cfg *config, args []string) error {
	if len(args) != 0 {
		fmt.Println("usage: pokedex   # takes no arguments")
		return nil
	}
	if len(cfg.pokedex) == 0 {
		fmt.Println("Your Pokedex is empty. Go catch some Pokémon!")
		return nil
	}

	fmt.Println("Your Pokedex:")
	var names []string
	for n := range cfg.pokedex {
		names = append(names, n)
	}
	sort.Strings(names)

	for _, n := range names {
		fmt.Printf(" - %s\n", n)
	}
	return nil
}
