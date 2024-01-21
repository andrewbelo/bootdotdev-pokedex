package main

import "fmt"

const LIMIT = 20

func commandMap(config *Config, _ []string) error {
	err := commandMapCore(config)
	if err != nil {
		return err
	}
	config.PokeClient.LocationAreaOffset += LIMIT
	return nil
}

func commandMapBack(config *Config, _ []string) error {
	if config.PokeClient.LocationAreaOffset < LIMIT {
		fmt.Println("You are already at the beginning of the list")
		return nil
	}
	config.PokeClient.LocationAreaOffset -= LIMIT
	return commandMapCore(config)
}

func commandMapCore(config *Config) error {
	areas, err := config.PokeClient.GetLocationAreas(LIMIT)
	if err != nil {
		return err
	}
	for _, area := range areas {
		fmt.Println(area.Name)
	}
	return nil
}
