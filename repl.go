package main

import (
	"strings"

	"github.com/JonahLargen/Pokedex/internal/pokeapi"
)

func cleanInput(text string) []string {
	inputs := []string{}
	lines := strings.Fields(text)

	for _, line := range lines {
		line = strings.ToLower(line)

		if line != "" {
			inputs = append(inputs, line)
		}
	}

	return inputs
}

type cliCommand struct {
	name        string
	description string
	usage       string
	callback    func(*config) error
}

type config struct {
	NextLocationAreaUrl     string
	PreviousLocationAreaUrl string
	CommandArgs             []string
	CaughtPokemon           map[string]pokeapi.PokemonResponse
}
