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
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Display this help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next map of the Pokedex",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous map of the Pokedex",
			callback:    commandMapBack,
		},
		"explore": {
			name:        "explore",
			description: "Explore a specific location area",
			callback:    commandExplore,
		},
	}
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
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
