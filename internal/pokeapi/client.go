package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"pokedexcli/internal/pokecache"
	"time"
)

type Client struct {
	LocationAreaOffset int
	Cache              *pokecache.Cache
}

func NewClient() *Client {
	return &Client{
		LocationAreaOffset: 0,
		Cache:              pokecache.NewCache(time.Duration(5) * time.Minute),
	}
}

func (c *Client) GetLocationAreas(limit int) ([]LocationArea, error) {
	var err error
	var body []byte
	locationAreaKey := fmt.Sprintf("location-area-%d-%d", c.LocationAreaOffset, limit)
	body, ok := c.Cache.Get(locationAreaKey)
	if !ok {
		body, err = c.MakeRequest(fmt.Sprintf(
			"location-area?offset=%d&limit=%d", c.LocationAreaOffset, limit,
		))
		if err != nil {
			return nil, err
		}
		c.Cache.Add(locationAreaKey, body)
	}

	var locAreas LocationAreaResponse
	err = json.Unmarshal(body, &locAreas)
	if err != nil {
		return nil, err
	}
	return locAreas.Results, nil
}

func (c *Client) GetPokemonDetails(name string) (PokemonDetailsResponse, error) {
	var err error
	var body []byte
	var pokemonDetails PokemonDetailsResponse
	pokemonDetailsKey := fmt.Sprintf("pokemon-details-%s", name)

	body, ok := c.Cache.Get(pokemonDetailsKey)
	if !ok {
		body, err = c.MakeRequest(fmt.Sprintf("pokemon/%s", name))
		if err != nil {
			return pokemonDetails, err
		}
		c.Cache.Add(pokemonDetailsKey, body)
	}

	err = json.Unmarshal(body, &pokemonDetails)
	if err != nil {
		return pokemonDetails, err
	}
	return pokemonDetails, nil
}

func (c *Client) GetLocationAreaDetails(name string) (LocationAreaDetailsResponse, error) {
	var err error
	var body []byte
	var locAreaDetails LocationAreaDetailsResponse
	locationAreaDetailsKey := fmt.Sprintf("location-area-details-%s", name)

	body, ok := c.Cache.Get(locationAreaDetailsKey)
	if !ok {
		body, err = c.MakeRequest(fmt.Sprintf("location-area/%s", name))
		if err != nil {
			return locAreaDetails, err
		}
		c.Cache.Add(locationAreaDetailsKey, body)
	}

	err = json.Unmarshal(body, &locAreaDetails)
	if err != nil {
		return locAreaDetails, err
	}
	return locAreaDetails, nil
}

func (c *Client) MakeRequest(suffix string) ([]byte, error) {
	resp, err := http.Get(API_ENDPINT + suffix)
	log.Println(API_ENDPINT + suffix)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(resp.Body)
	resp.Body.Close()
	if resp.StatusCode > 299 {
		return nil, errors.New(fmt.Sprintf(
			"Response failed with status code: %d and\nbody: %s\n",
			resp.StatusCode,
			body,
		))
	}
	if err != nil {
		return nil, err
	}
	return body, nil
}
