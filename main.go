package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexcli/internal/pokeapi"
	"strings"
)

type PokemonStats struct {
	BaseStat int
	Effort   int
	Name     string
	URL      string
}

type PokemonType struct {
	Slot int
	Name string
	URL  string
}

type Pokemon struct {
	Name           string
	BaseExperience int
	Height         int
	IsDefault      bool
	Order          int
	Weight         int
	Stats          []PokemonStats
	Types          []PokemonType
}

type Config struct {
	PokeClient *pokeapi.Client
	Pokedex    map[string]Pokemon
}

func main() {
	config := Config{
		PokeClient: pokeapi.NewClient(),
		Pokedex:    make(map[string]Pokemon),
	}
	reader := bufio.NewScanner(os.Stdin)
	for true {
		cmds := getCommands()

		fmt.Print("pokedex > ")
		reader.Scan()
		input_words := cleanInput(reader.Text())
		if len(input_words) == 0 {
			continue
		}
		cmd := input_words[0]
		command, ok := cmds[cmd]
		if !ok {
			fmt.Println("Command not found")
			printOptions()
			continue
		}
		error := command.callback(&config, input_words[1:])
		if error != nil {
			fmt.Println(error)
			break
		}
	}
}

func cleanInput(input string) []string {
	return strings.Fields(strings.ToLower(input))
}
