package main

import (
	"errors"
	"fmt"
	"sort"
)

func commandExplore(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: explore <location-area>")
	}
	areaName := args[0]
	fmt.Printf("Exploring %s... \n", areaName)

	area, err := cfg.pokeapiClient.GetLocationArea(areaName)
	if err != nil {
		return err
	}
	var names []string

	for _, enc := range area.PokemonEncounters {
		names = append(names, enc.Pokemon.Name)
	}
	sort.Strings(names)

	fmt.Println("Found Pokemon:")
	for _, n := range names {
		fmt.Printf(" - %s\n", n)
	}

	return nil
}
