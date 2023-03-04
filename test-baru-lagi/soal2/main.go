package main

import (
	"fmt"
)

func reverseString(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < len(runes)/2; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	input := "jatis"
	output := reverseString(input)
	fmt.Println(output)
}
