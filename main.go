package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/babanini95/pokedexcli/internal"
)

func main() {
	commands = generateCommand()
	scanner := bufio.NewScanner(os.Stdin)
	conf.Cache = internal.NewCache(5 * time.Second)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		cleanedInput := cleanInput(scanner.Text())

		if len(cleanedInput) == 0 {
			continue
		}

		userCommand := cleanedInput[0]

		if cmd, ok := commands[userCommand]; !ok {
			fmt.Println("Unknown command")
		} else {
			err := cmd.callback(&conf)
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
