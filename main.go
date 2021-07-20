package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]

	for _, word := range args {
		letters := strings.Split(word, "")

		for _, letter := range letters {
			fmt.Print(":alphabet-white-" + letter + ":")
		}

		fmt.Print("  ")
	}
}
