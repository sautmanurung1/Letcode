package main

import (
	"fmt"
)

func lastLetters(word string) string {
	// First Reverse the string
	var reversedWord string
	for i := len(word) - 1; i >= 0; i-- {
		reversedWord += string(word[i])
	}
	// Second Delete the last letter of the reversed string
	var lastLetter string
	for i := 0; i < len(reversedWord)-1; i++ {
		lastLetter += string(reversedWord[i])
	}
	// Third get only two first letters of the reversed string
	var firstLetters string
	for i := 0; i < 2; i++ {
		firstLetters += string(reversedWord[i])
	}
	// Fourth Give Space the first Letter
	var firstLattersWithSpace string
	firstLattersWithSpace += string(firstLetters[0])
	for i := 1; i < len(firstLetters); i++ {
		firstLattersWithSpace += " "
		firstLattersWithSpace += string(firstLetters[i])
	}
	return firstLattersWithSpace
}

func main() {
	fmt.Println(lastLetters("APPLE"))  // Output : E L
	fmt.Println(lastLetters("ORANGE")) // Output : E G
}
