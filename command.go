package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/babanini95/pokedexcli/internal"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

type config struct {
	Next  *string
	Prev  *string
	Cache *internal.Cache
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
		"explore": {
			name:        "explore",
			description: "Lists all the pokemon located in a given area",
			callback:    commandExplore,
		},
	}
}

func commandExit(c *config, args []string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *config, args []string) error {
	output := "Welcome to the Pokedex!\nUsage:\n"

	for _, command := range commands {
		output += fmt.Sprintf("\n%s: %s", command.name, command.description)
	}

	fmt.Println(output)
	return nil
}

func commandMap(c *config, args []string) error {
	url := "https://pokeapi.co/api/v2/location-area/"
	if c.Next != nil {
		url = *c.Next
	}
	err := getDataAreaResult(url, c)
	return err
}

func commandMapB(c *config, args []string) error {
	if c.Prev == nil {
		fmt.Println("you're on the first page")
		return nil
	}
	err := getDataAreaResult(*c.Prev, c)
	return err
}

func commandExplore(c *config, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("unknown command")
	}
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	url := baseUrl + args[0]
	var areaData internal.AreaData
	var err error
	fmt.Println("url: " + url)
	data, ok := c.Cache.Get(url)
	if ok {
		err = json.Unmarshal(data, &areaData)
		if err != nil {
			return err
		}
	} else {
		areaData, err = internal.GetEncounteredPokemon(url)
		if err != nil {
			return err
		}

		cacheValue, err := json.Marshal(areaData)
		if err != nil {
			return err
		}

		// Both name and Id return the same data
		urlWithName := baseUrl + areaData.Name
		urlWithId := fmt.Sprintf("%s%d", baseUrl, areaData.ID)
		c.Cache.Add(urlWithName, cacheValue)
		c.Cache.Add(urlWithId, cacheValue)
	}

	fmt.Printf("Exploring %s...\nFound Pokemon:\n", areaData.Name)
	for _, pokemon := range areaData.PokemonEncounters {
		fmt.Println(" - " + pokemon.Pokemon.Name)
	}
	return nil
}

func getDataAreaResult(url string, c *config) error {
	var dataArea internal.LocationAreas
	var err error

	data, ok := c.Cache.Get(url)

	if ok {
		err = json.Unmarshal(data, &dataArea)
		if err != nil {
			return err
		}
	} else {
		dataArea, err = internal.GetLocations(url)
		if err != nil {
			return err
		}

		cacheValue, err := json.Marshal(dataArea)
		if err != nil {
			return err
		}

		c.Cache.Add(url, cacheValue)
	}

	c.Next = dataArea.Next
	c.Prev = dataArea.Previous
	for _, result := range dataArea.Results {
		fmt.Println(result.Name)
	}
	return nil
}
