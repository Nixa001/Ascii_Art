package main

import (
	"os"
)

func asciiArtOutput(finalText string) {
	result := os.Args[1][9:]
	file, err := os.Create(result)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	_, err = file.WriteString(finalText)
	if err != nil {
		panic(err)
	}
}