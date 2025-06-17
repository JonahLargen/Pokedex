package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  Go is   great!  ",
			expected: []string{"go", "is", "great!"},
		},
		{
			input:    "  Multiple   spaces   here  ",
			expected: []string{"multiple", "spaces", "here"},
		},
		{input: "  Leading and trailing spaces  ",
			expected: []string{"leading", "and", "trailing", "spaces"},
		},
		{
			input:    "  \t\n  \r  ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected %d words but got %d", c.input, len(c.expected), len(actual))
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("For input '%s', expected '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}
