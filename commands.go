package main

import "fmt"

type cliCommand struct {
	name        string
	description string
	callback    func(config *Config, args []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the names of the next 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of the Pokemon that can be found in the given location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect",
			description: "Displays the details of a Pokemon that you have caught",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Displays the names of the Pokemon that you have caught",
			callback:    commandPokedex,
		},
	}
}

func printOptions() {
	cmdTexts := getCommands()

	fmt.Println("\nUsage:")
	for _, cmd := range cmdTexts {
		fmt.Printf("\t%s - %s\n", cmd.name, cmd.description)
	}
}
