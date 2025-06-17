package main

import (
	"strings"
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
	callback    func() error
}

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
			callback:    help,
		},
	}
}
