package main

import (
	"bufio"
	"fmt"
	"github.com/NicholasRodrigues/pokedexcli/internal/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient    pokeapi.Client
	nextLocationsUrl *string
	prevLocationsUrl *string
	caughtPokemon    map[string]pokeapi.Pokemon
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()

		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandName := words[0]

		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg, words[1:]...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
