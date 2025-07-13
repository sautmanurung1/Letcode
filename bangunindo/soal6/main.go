package main

import "strings"

func isVowel(char byte) bool {
	vowels := "aeiouAEIOU"
	return strings.ContainsRune(vowels, rune(char))
}

func vowelsubstring(s string) int {
	count := 0
	vowelCount := 0
	window := make(map[byte]int)

	for right := 0; right < len(s); right++ {
		if isVowel(s[right]) {
			if window[s[right]] == 0 {
				vowelCount++
			}
			window[s[right]]++
		}

		for vowelCount == 5 {
			if vowelCount == 5 {
				count += len(s) - right
			}
			if isVowel(s[right]) {
				window[s[right]]--
				if window[s[right]] == 0 {
					vowelCount--
				}
			}
		}
	}

	return count
}
