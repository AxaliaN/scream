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

        for _, letter := range letters {
            fmt.Print(":alphabet-white-" + letter + ":")
        }

        if i != (len(args) - 1) {
            fmt.Print("  ")
        }
    }
}
