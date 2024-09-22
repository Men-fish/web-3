package main

import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Scanln(&input)
	chars := strings.Split(input, "")
	result := strings.Join(chars, "*")
	fmt.Println(result)
}
