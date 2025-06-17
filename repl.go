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
