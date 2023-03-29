package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var HeightChar = 9
var tabChars = map[int][]string{}

func main() {
	RecupAsciiArt()

	if len(os.Args) < 2 {
		fmt.Println("Invalid Command")
		fmt.Println("Usage: go run . \"input string\"")
		os.Exit(0)
	}

	input := os.Args[1]
	asciiArt := generateAsciiArt(input)
	fmt.Print(asciiArt)
}

func RecupAsciiArt() {
	data, _ := ioutil.ReadFile("standard.txt")

	lines := strings.Split(string(data), "\n")
	characters := len(lines) / HeightChar
	for i := 0; i < characters; i++ {
		charLines := lines[i*HeightChar : (i+1)*HeightChar]
		tabChars[int(i)] = charLines

	}
}

func generateAsciiArt(input string) string {
	var result string
	for i := 1; i < HeightChar; i++ {
		for _, char := range input {
			if int(char) < 32 {
				fmt.Println("Invalid Input")
			}
			chars := int(char) - 32
			line := tabChars[chars][i]
			result += string(line)
		}
		result += "\n"
	}
	return result
}
