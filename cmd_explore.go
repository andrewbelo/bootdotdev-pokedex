package main

import (
	"fmt"
)

func commandExplore(config *Config, args []string) error {
	if len(args) != 1 {
		fmt.Errorf("Usage: explore <location>")
		return nil
	}
	fmt.Printf("Exploring %s...\n", args[0])
	location := args[0]
	locationDetails, err := config.PokeClient.GetLocationAreaDetails(location)
	if err != nil {
		return err
	}
	fmt.Println("Found the following Pokemon in the area:")
	for _, encounter := range locationDetails.PokemonEncounters {
		fmt.Printf("  - %s\n", encounter.Pokemon.Name)
	}
	return nil
}
