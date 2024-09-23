package main

import (
	"fmt"
)

func main() {
	var s string
	fmt.Scan(&s)

	maxDigit := '0'
	for _, char := range s {
		if char > maxDigit {
			maxDigit = char
		}
	}

	fmt.Printf("%c\n", maxDigit)
}
