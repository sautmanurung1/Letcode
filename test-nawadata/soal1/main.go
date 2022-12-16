package main

import (
	"fmt"
	"strings"
)

func sortAlphabet(s string) {
	s = strings.ToLower(s)
	vokalString := ""
	konsonanString := ""
	vokal := map[string]string{
		"a": "a",
		"i": "i",
		"u": "u",
		"e": "e",
		"o": "o",
	}
	for _, ch := range s {
		if vokal[string(ch)] != "" {
			vokalString += string(ch)
		} else if vokal[string(ch)] == "" && string(ch) != " " {
			konsonanString += string(ch)
		}
	}
	fmt.Println("Vowel Characters:")
	fmt.Println(vokalString)
	fmt.Println("Consonant Characters:")
	fmt.Println(konsonanString)
}

func main() {
	sortAlphabet("Sample Case")
	sortAlphabet("Next Case")
}
