package main

import (
    "fmt"
	"os"
)

func main() {
	input := os.Args[1]

	for _,e := range(input) {
		fmt.Printf("Rune: %c, Unicode: %U, Integer: %d\n", e, e, e)
	}
}

