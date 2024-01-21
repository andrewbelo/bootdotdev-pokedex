package main

import "fmt"

func commandInspect(config *Config, args []string) error {
	if len(args) != 1 {
		fmt.Errorf("Usage: inspect <pokemon>")
		return nil
	}
	pokemon, ok := config.Pokedex[args[0]]
	if !ok {
		fmt.Printf("You haven't caught %s yet!\n", args[0])
		return nil
	}
	fmt.Printf("Inspecting %s...\n", args[0])
	fmt.Printf("Name: %s\n", pokemon.Name)
	fmt.Printf("Height: %d\n", pokemon.Height)
	fmt.Printf("Weight: %d\n", pokemon.Weight)
	fmt.Printf("Order: %d\n", pokemon.Order)
	fmt.Println("Stats:")
	for _, stat := range pokemon.Stats {
		fmt.Printf("  - %s: %d\n", stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, type_ := range pokemon.Types {
		fmt.Printf("  - %s\n", type_.Name)
	}
	return nil
}
