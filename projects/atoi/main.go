package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input int
	fmt.Scan(&input)

	result := ""
	for _, digit := range strconv.Itoa(input) {
		num, _ := strconv.Atoi(string(digit))
		result += strconv.Itoa(num * num)
	}

	fmt.Println(result)
}
