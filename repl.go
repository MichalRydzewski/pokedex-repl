package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func CleanInput(text string) []string {
	sliceOfWords := strings.Fields(text)
	return sliceOfWords
}

func exit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func help() error {
	fmt.Printf("Welcome to the Pokedex!\n")
	fmt.Printf("Usage:\nhelp: Displays a help message\nexit: Exit the Pokedex\n")
	return nil
}

var commands = map[string]cliCommand{
	"exit": {
		name:        "exit",
		description: "Exit the Pokedex",
		callback:    exit,
	},
	"help": {
		name:        "help",
		description: "Provide help",
		callback:    help,
	},
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		word := strings.ToLower(scanner.Text())
		cmd, ok := commands[word]
		if !ok {
			fmt.Printf("Unknown command")
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Println(err)
		}
	}
}
