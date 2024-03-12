package main

import (
	"github.com/NicholasRodrigues/pokedexcli/internal/pokeapi"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  ",
			expected: []string{},
		},
		{
			input:    "  hello  ",
			expected: []string{"hello"},
		},
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "  HellO  World  ",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
}

func TestGetCommands(t *testing.T) {
	commands := getCommands()
	if len(commands) != 4 {
		t.Errorf("expected 4 commands, got %v", len(commands))
	}
}

func TestCommandExit(t *testing.T) {
	cfg := &config{}
	err := commandExit(cfg)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestCommandHelp(t *testing.T) {
	cfg := &config{}
	err := commandHelp(cfg)
	if err != nil {
		t.Errorf("expected nil, got %v", err)
	}
}

func TestCommandMapF(t *testing.T) {
	cfg := &config{
		pokeapiClient: pokeapi.Client{},
	}
	err := commandMapF(cfg)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}

func TestCommandMapB(t *testing.T) {
	cfg := &config{
		pokeapiClient: pokeapi.Client{},
	}
	err := commandMapB(cfg)
	if err == nil {
		t.Errorf("expected error, got nil")
	}
}
