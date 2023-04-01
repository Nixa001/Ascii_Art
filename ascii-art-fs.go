package main

import (
	"os"
)

func asciiArtFs() string {
	args := os.Args
	value := "banner/standard.txt"
	if len(args) == 4 {
		value = "banner/" + os.Args[3] + ".txt"
		return value
	}
	return value
}
