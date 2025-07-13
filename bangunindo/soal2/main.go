package main

import (
	"fmt"
	"strconv"
)

func changeAds(base10 int32) int32 {
	// Convert base10 to binary string
	binaryStr := strconv.FormatInt(int64(base10), 2)

	// Negate each bit in the binary string
	var negatedBinaryStr string
	for _, bit := range binaryStr {
		if bit == '0' {
			negatedBinaryStr += "1"
		} else {
			negatedBinaryStr += "0"
		}
	}

	// Convert negated binary string back to base 10
	negatedBase10, _ := strconv.ParseInt(negatedBinaryStr, 2, 32)

	return int32(negatedBase10)
}

func main() {
	// Sample cases
	fmt.Println(changeAds(50))  // Output: 13
	fmt.Println(changeAds(100)) // Output: 27
}
