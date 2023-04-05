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
	arg1 := os.Args[1]
	if len(args) < 2 || len(args) > 3 {
		fmt.Println("Usage: go run . [STRING] [BANNER]")
		fmt.Println("\nEX: go run . something standard")
		os.Exit(0)
	}
	ReadBanner()
	GenerateAscii(arg1)
}

func ReadBanner() {
	arg2 := ReadBannerTextArg()
	var tabChars []string
	scanner, err := ioutil.ReadFile(arg2)
	if err != nil {
		fmt.Println("Invalid banner")
		os.Exit(1)
	}
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

func GenerateAscii(input string) {
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
