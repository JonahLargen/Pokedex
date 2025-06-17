package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		if scanner.Scan() {
			input := scanner.Text()
			cleanedInput := cleanInput(input)
			if len(cleanedInput) == 0 {
				fmt.Println("No input provided. Please enter a command.")
				continue
			}
			firstWord := cleanedInput[0]
			fmt.Printf("Your command was: %v\n", firstWord)
		} else {
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading input: %v\n", err)
			}
			break
		}
	}
}
