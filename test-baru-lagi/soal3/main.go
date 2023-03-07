package main

import (
	"fmt"
	"strings"
)

func removeSpaces(inputStr string) string {
	return strings.ReplaceAll(inputStr, " ", "")
}

func changeToLowercase(input string) string {
	return strings.ToLower(input)
}

func countCharacters(inputStr string) map[string]int {
	pecahtext := make(map[string]int)
	for _, char := range inputStr {
		charStr := string(char)
		if _, ok := pecahtext[charStr]; ok {
			pecahtext[charStr] += 1
		} else {
			pecahtext[charStr] = 1
		}
	}
	return pecahtext
}

func printCharacterCounts(inputStr string) string {
	counts := countCharacters(inputStr)
	var outputStr string
	for char, count := range counts {
		if count == 1 {
			outputStr += char
		} else {
			outputStr += fmt.Sprintf("%d%s", count, char)
		}
	}
	return outputStr
}

func main() {
	input := "dani Maulana"
	fmt.Println(printCharacterCounts(removeSpaces(changeToLowercase(input))))
}
