package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input one line of words: ")
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(text)
	text = strings.ReplaceAll(text, " ", "")
	vowel := "aiueo"
	resultVowel := ""
	resultConsonant := ""
	for _, ch := range text {
		if strings.ContainsRune(vowel, ch) {
			resultVowel += string(ch)
		} else {
			resultConsonant += string(ch)
		}
	}
	fmt.Println("Vowel Characters:")
	fmt.Println(resultVowel)
	fmt.Println("Consonant Characters:")
	fmt.Println(resultConsonant)
}
