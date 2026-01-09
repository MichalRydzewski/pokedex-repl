package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	sliceOfWords := strings.Fields(text)
	return sliceOfWords
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		word := strings.ToLower(scanner.Text())
		fmt.Printf("Your command was: %v\n", word)
	}
}
