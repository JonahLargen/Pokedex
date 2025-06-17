package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	commands := getCommands()

	for {
		fmt.Print("Pokedex> ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) == 0 {
				fmt.Println("No input provided. Please enter a command.")
				continue
			}
			commandName := cleanedInput[0]
			command, ok := commands[commandName]
			if !ok {
				fmt.Println("Unknown command:", commandName)
			} else {
				err := command.callback()
				if err != nil {
					fmt.Printf("Error executing command '%s': %v\n", commandName, err)
				}
			}
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			break
		}
	}
}
