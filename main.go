package main

import (
    "fmt"
    "os"
    "strings"
)

func main() {
    args := os.Args[1:]

    for i := 0; i < len(args); i++ {
        letters := strings.Split(args[i], "")

        // Loop over the letters in the word
        for _, letter := range letters {
            fmt.Print(":alphabet-white-" + letter + ":")
        }

        // If we are not at the final word, add a couple of spaces
        if i != (len(args) - 1) {
            fmt.Print("  ")
        }
    }

    // Print an empty line for easy copying in the terminal
    fmt.Println()
}
