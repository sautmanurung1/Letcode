package main

import "fmt"

const MAP_LENGTH = 256

func lengthOfLongestSubstring(substring string) int {
	currentSubstringHashMap := [MAP_LENGTH]bool{}
	currentSubstringHashMapPosition := [MAP_LENGTH]int{}
	sequenceLengthMax := 0
	sequenceLengthCurrent := 0
	substringLength := len(substring)
	for currentSymbolIndex := 0; currentSymbolIndex < substringLength; currentSymbolIndex++ {
		symbol := substring[currentSymbolIndex]
		isSymbolExists := currentSubstringHashMap[symbol]

		if isSymbolExists {
			if sequenceLengthCurrent > sequenceLengthMax {
				sequenceLengthMax = sequenceLengthCurrent
			}
			currentSymbolIndex = currentSubstringHashMapPosition[symbol]

			sequenceLengthCurrent = 0
			currentSubstringHashMap = [MAP_LENGTH]bool{}
			continue
		}
		currentSubstringHashMap[symbol] = true
		currentSubstringHashMapPosition[symbol] = currentSymbolIndex

		sequenceLengthCurrent++
		if sequenceLengthCurrent > sequenceLengthMax {
			sequenceLengthMax = sequenceLengthCurrent
		}
	}
	return sequenceLengthMax
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abbbbbcccccdddddeeeeffffgggg"))
}
