package main

import (
	"fmt"
	"strings"
)

func CleanInput(text string) []string {
	sliceOfWords := strings.Fields(text)

	return sliceOfWords
}

func main() {
	fmt.Println("Hello, World!")
}
