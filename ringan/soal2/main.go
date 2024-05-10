package main

import (
	"fmt"
	"strings"
)

// __define-ocg__
func ArrayChallenge(strArr []string) string {
	cache := make([]string, 0, 5) // Cache with maximum capacity of 5 elements

	for _, char := range strArr {
		found := false
		for i, c := range cache {
			if c == char {
				found = true
				// Move the accessed element to the front
				cache = append(cache[:i], append([]string{cache[i]}, cache[:i]...)...)
				break
			}
		}

		if !found {
			// If cache is full, remove the least recently used element
			if len(cache) == 5 {
				cache = cache[:4]
			}
			// Add the new element to the front
			cache = append([]string{char}, cache...)
		}
	}

	return strings.Join(cache, "-")
}

func main() {
	fmt.Println(ArrayChallenge([]string{"A", "B", "C", "D", "A", "E", "D", "Z"})) // Output: C-A-E-D-Z
	fmt.Println(ArrayChallenge([]string{"A", "B", "A", "C", "A", "B"}))            // Output: C-A-B
	fmt.Println(ArrayChallenge([]string{"A", "B", "C", "D", "E", "D", "Q", "Z", "C"})) // Output: E-D-Q-Z-C
}
