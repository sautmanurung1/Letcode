package main

import (
	"fmt"
	"sort"
)

func countChars(str string) map[rune]int {
	counts := make(map[rune]int)
	for _, ch := range str {
		if ch == ' ' {
			continue
		}
		counts[ch]++
	}
	return counts
}

func main() {
	str := "random string"
	counts := countChars(str)

	// Convert the map to a slice of pairs
	pairs := make([][2]interface{}, 0, len(counts))
	for k, v := range counts {
		pairs = append(pairs, [2]interface{}{k, v})
	}

	// Sort the pairs by the UTF-8 value of the character
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0].(rune) < pairs[j][0].(rune)
	})

	// Print the counts in sorted order
	for _, pair := range pairs {
		ch := pair[0].(rune)
		count := pair[1].(int)
		fmt.Printf("%c: %d\n", ch, count)
	}
}
