package main

import "fmt"

func commandPokedex(config *Config, _ []string) error {
	if len(config.Pokedex) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}
	fmt.Println("Here is a list of all the Pokemon in the Pokedex:")
	for _, pokemon := range config.Pokedex {
		fmt.Printf("  - %s\n", pokemon.Name)
	}
	return nil
}
