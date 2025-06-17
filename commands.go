package main

import (
	"fmt"
	"os"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	commands := getCommands()
	for _, cmd := range commands {
		fmt.Printf("\n%s: %s", cmd.name, cmd.description)
	}
	fmt.Println()
	return nil
}
