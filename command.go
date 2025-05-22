package main

import (
	"fmt"
	"os"

	"github.com/babanini95/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next *string
	Prev *string
}

var commands map[string]cliCommand
var conf config

func generateCommand() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the next 20 location areas in Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous 20 location areas in Pokemon world",
			callback:    commandMapB,
		},
	}
}

func commandExit(c *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config) error {
	output := "Welcome to the Pokedex!\nUsage:\n"

	for _, command := range commands {
		output += fmt.Sprintf("\n%s: %s", command.name, command.description)
	}

	fmt.Println(output)
	return nil
}

func commandMap(c *config) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != nil {
		url = *c.Next
	}
	err := getDataAreaResult(url, c)
	return err
}

func commandMapB(c *config) error {
	if c.Prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	err := getDataAreaResult(*c.Prev, c)
	return err
}

func getDataAreaResult(url string, c *config) error {
	dataArea, err := internal.GetLocations(url)
	if err != nil {
		return err
	}
	c.Next = dataArea.Next
	c.Prev = dataArea.Previous
	for _, result := range dataArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}
