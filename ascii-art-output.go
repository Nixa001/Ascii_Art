package main

import (
	"fmt"
	"os"
	"strings"
)

func asciiArtOutput(finalText string) {
	var result string
	if len(os.Args[1])>9 && strings.HasPrefix(os.Args[1], "--output="){
		result = os.Args[1][9:]
	}
	file, err := os.Create(result)
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
		fmt.Println("\nEX: go run . --output=<fileName.txt> something standard")
		os.Exit(1)
	}
	defer file.Close()
	_, err = file.WriteString(finalText)
	if err != nil {
		fmt.Println("Error")
	}
}