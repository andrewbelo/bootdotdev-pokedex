package main

import (
	"fmt"
	"math/rand"
	"pokedexcli/internal/pokeapi"
)

func commandCatch(config *Config, args []string) error {
	if len(args) != 1 {
		fmt.Errorf("Usage: catch <pokemon>")
		return nil
	}
	fmt.Printf("Trying to catch %s...\n", args[0])
	caught := rand.Intn(2) == 0
	if !caught {
		fmt.Printf("Oh no! %s got away!\n", args[0])
		return nil
	}
	fmt.Printf("Caught %s...\n", args[0])
	pokemon := args[0]
	pokemonDetails, err := config.PokeClient.GetPokemonDetails(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf(
		"You caught %s!\nAdding it to your Pokedex...\n",
		pokemonDetails.Name,
	)
	config.Pokedex[pokemonDetails.Name] = Pokemon{
		Name:           pokemonDetails.Name,
		BaseExperience: pokemonDetails.BaseExperience,
		Height:         pokemonDetails.Height,
		IsDefault:      pokemonDetails.IsDefault,
		Order:          pokemonDetails.Order,
		Weight:         pokemonDetails.Weight,
		Stats:          parseStats(pokemonDetails),
		Types:          parseTypes(pokemonDetails),
	}
	return nil
}

func parseTypes(pokemon_details pokeapi.PokemonDetailsResponse) []PokemonType {
	types := pokemon_details.Types
	parsedTypes := make([]PokemonType, len(types))
	for i, type_ := range types {
		parsedTypes[i] = PokemonType{
			Slot: type_.Slot,
			Name: type_.Type.Name,
			URL:  type_.Type.URL,
		}
	}
	return parsedTypes
}

func parseStats(pokemon_details pokeapi.PokemonDetailsResponse) []PokemonStats {
	stats := pokemon_details.Stats
	parsedStats := make([]PokemonStats, len(stats))
	for i, stat := range stats {
		parsedStats[i] = PokemonStats{
			BaseStat: stat.BaseStat,
			Effort:   stat.Effort,
			Name:     stat.Stat.Name,
			URL:      stat.Stat.URL,
		}
	}
	return parsedStats
}
