package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	var userInput string
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		userInput = scanner.Text()
		fmt.Printf("Your command was: %s\n", cleanInput(userInput)[0])
		userInput = ""
	}
}

func cleanInput(text string) []string {
	cleanText := strings.Fields(strings.ToLower(text))
	return cleanText
}
