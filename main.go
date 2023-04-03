package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var HeightChar = 9
var tabChar = map[int][]string{}

func main() {
	RecupAsciiArt()

	if len(os.Args) != 2 {
		fmt.Println("Invalid Command")
		fmt.Println("Usage: go run . \"Input\"")
		os.Exit(0)
	}

	input := os.Args[1]
	asciiArt := generateAsciiArt(input)
	fmt.Print(asciiArt)
}

func RecupAsciiArt() {
	var tabChars []string
	scanner, _ := ioutil.ReadFile("banner/standard.txt")
	data := bufio.NewScanner(strings.NewReader(string(scanner)))
	for data.Scan() {
		lines := data.Text()
		tabChars = append(tabChars, lines)
	}
	characters := len(tabChars) / HeightChar
	for i := 0; i < characters; i++ {
		charLines := tabChars[i*HeightChar : (i+1)*HeightChar]
		tabChar[int(i)] = charLines

	}
}

func generateAsciiArt(input string) string {
	var result string
	inputLines := strings.Split(input, "\\n")
	for _, inputLine := range inputLines {
		for i := 1; i < HeightChar; i++ {
			for _, char := range inputLine {
				if int(char) < 32 {
					fmt.Println("Invalid Input")
				}
				chars := int(char) - 32
				line := tabChar[chars][i]
				result += string(line)
			}
			result += "\n"
		}
	}
	return result
}
