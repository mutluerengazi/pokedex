package main

import (
	"errors"
	"fmt"
	"sort"
)

func commandInspect(cfg *config, args []string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}
	name := args[0]

	pkm, ok := cfg.pokedex[name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %s\n", pkm.Name)
	fmt.Printf("Height: %d\n", pkm.Height)
	fmt.Printf("Weight: %d\n", pkm.Weight)

	fmt.Println("Stats:")
	for _, s := range pkm.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.Base)
	}

	// gather and sort type names for nice output
	var tnames []string
	for _, t := range pkm.Types {
		tnames = append(tnames, t.Type.Name)
	}
	sort.Strings(tnames)

	fmt.Println("Types:")
	for _, tn := range tnames {
		fmt.Printf("  - %s\n", tn)
	}
	return nil

}
