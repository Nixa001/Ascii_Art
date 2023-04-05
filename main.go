package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var tabChar = map[int][]string{}
var HeightChar = 9

func main() {
	args := os.Args
	arg1 := os.Args[2]
	if len(args) < 3 || len(args) > 4 {
		fmt.Println("Invalid command")
		fmt.Println("Usage: go run . \"Input string\" [Banner]")
		os.Exit(0)
	}
	ReadAsciiTab()
	output := GenerateAscii(arg1)
	asciiArtOutput(output)
}
func ReadAsciiTab() {
	arg2 := asciiArtFs()
	var tabChars []string
	scanner, _ := ioutil.ReadFile(arg2)
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
func GenerateAscii(input string) string {
	var output string
	divInput := strings.Split(input, "\\n")
	for _, InputPart := range divInput {

		for i := 1; i < HeightChar; i++ {
			for _, runeValue := range InputPart {
				if int(runeValue) >= 32 && int(runeValue) <= 126 {
					NumCharsInTab := int(runeValue) - 32
					outputLine := tabChar[NumCharsInTab][i]
					if outputLine != "" {
						output += outputLine
					}

				} else {
					fmt.Println("Invalid Input")
					os.Exit(0)
				}
			}
			if InputPart != "" {
				output += "\n"

			}
		}
	}
	return output
}
