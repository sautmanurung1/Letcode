package main

import (
	"fmt"
	"math"
)

func isPrime(number int) bool {
	start := 2
	for start <= int(math.Sqrt(float64(number))) {
		if number%start < 1 {
			return false
		}
		start++
	}
	return number > 1
}

func findPrimeByRange(start int, finish int) []int {
	realStart := start
	realEnd := finish
	if start > finish {
		realStart = finish
		realEnd = start
	}
	var numbers []int
	for current := realStart; current <= realEnd; current++ {
		if isPrime(current) {
			numbers = append(numbers, current)
		}
	}
	return numbers
}

func main() {
	fmt.Println(findPrimeByRange(11, 40))
}
