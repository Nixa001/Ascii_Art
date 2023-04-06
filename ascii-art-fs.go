package main

import (
	"os"
)

func ReadBannerTextArg() string {
	args := os.Args
	value := "banner/standard.txt"
	if len(args) == 3 {
		value = "banner/" + os.Args[2] + ".txt"
		return value
	}
	return value
}
