package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	commands = generateCommand()
	scanner := bufio.NewScanner(os.Stdin)

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
			err := cmd.callback()
			if err != nil {
				fmt.Println(err)
			}
		}
	}
}
