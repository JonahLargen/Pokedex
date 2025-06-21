package main

import (
	"fmt"
	"os"

	"github.com/JonahLargen/Pokedex/internal/pokeapi"
)

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			usage:       "",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display this help message",
			usage:       "",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next map of the Pokedex",
			usage:       "",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous map of the Pokedex",
			usage:       "",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location area",
			usage:       "explore <location-area>",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokémon by name",
			usage:       "catch <pokemon-name>",
			callback:    commandCatch,
		},
	}
}

var commandOrder = []string{
	"exit",
	"help",
	"map",
	"mapb",
	"explore",
	"catch",
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commands := getCommands()
	for _, name := range commandOrder {
		cmd := commands[name]
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
		if cmd.usage != "" {
			fmt.Printf("\n   Usage: %s", cmd.usage)
		}
	}
	fmt.Println()
	return nil
}

func commandMap(cfg *config) error {
	locations, err := pokeapi.GetLocationAreas(cfg.NextLocationAreaUrl)
	if err != nil {
		return err
	}
	cfg.NextLocationAreaUrl = locations.Next
	cfg.PreviousLocationAreaUrl = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapBack(cfg *config) error {
	locations, err := pokeapi.GetLocationAreas(cfg.PreviousLocationAreaUrl)
	if err != nil {
		return err
	}
	cfg.NextLocationAreaUrl = locations.Next
	cfg.PreviousLocationAreaUrl = locations.Previous
	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandExplore(cfg *config) error {
	if len(cfg.CommandArgs) < 1 {
		return fmt.Errorf("please provide a location area name; E.g. 'explore viridian-forest'")
	}
	areaName := cfg.CommandArgs[0]
	if areaName == "" {
		return fmt.Errorf("location area name cannot be empty")
	}
	resp, err := pokeapi.GetLocationAreaInfo(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", resp.Location.Name)
	for _, encounter := range resp.PokemonEncounters {
		fmt.Printf("- %v\n", encounter.Pokemon.Name)
	}
	return nil
}

func commandCatch(cfg *config) error {
	if len(cfg.CommandArgs) < 1 {
		return fmt.Errorf("please provide a Pokémon name to catch; E.g. 'catch pikachu'")
	}
	pokemonName := cfg.CommandArgs[0]
	if pokemonName == "" {
		return fmt.Errorf("pokémon name cannot be empty")
	}
	resp, err := pokeapi.GetPokemonInfo(pokemonName)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", resp.Name)
	caught := pokeapi.AttemptCatch(resp)
	if !caught {
		fmt.Printf("%s escaped!\n", resp.Name)
		return nil
	}
	fmt.Printf("%s was caught!\n", resp.Name)
	if _, exists := cfg.CaughtPokemon[resp.Name]; !exists {
		cfg.CaughtPokemon[resp.Name] = *resp
	}
	return nil
}
