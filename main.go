package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var tabChars = map[int][]string{}
var taille = 9

func main() {
	args := os.Args
	arg1 := os.Args[1]
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Invalid command")
		fmt.Println("Usage: go run . \"Input string\" [Banner]")
		os.Exit(0)
	}
	ReadAsciiTab()
	GenerateAscii(arg1)
}
func ReadAsciiTab() {
	arg2 := Takebanner()
	data, _ := ioutil.ReadFile(arg2)
	dataChar := strings.Split(string(data), "\n")
	NumberChars := len(dataChar) / taille
	for i := 0; i < NumberChars; i++ {
		character := dataChar[i*taille : (i+1)*taille]
		tabChars[int(i)] = character
	}
}
func GenerateAscii(input string) {
	var output string
	divInput := strings.Split(input, "\\n")
	for _, InputPart := range divInput {

		for i := 1; i < taille; i++ {
			for _, runeValue := range InputPart {
				if int(runeValue) >= 32 && int(runeValue) <= 126 {
					NumCharsInTab := int(runeValue) - 32
					outputLine := tabChars[NumCharsInTab][i]
					if outputLine != "" {
						output += outputLine
					}

				} else {
					fmt.Println("Invalid Input")
					os.Exit(0)
				}
			}
			output += "\n"
		}
		if InputPart == "" {
			fmt.Println()
		} else {
			fmt.Print(output)
		}
		output = ""
	}

}
