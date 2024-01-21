package main

import "fmt"

func commandHelp(_ *Config, _ []string) error {
	fmt.Println("Welcome to the Pokedex!")
	printOptions()
	return nil
}
