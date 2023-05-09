package main

import "fmt"

func fibonacciSum(X, Y, N int) (int, int) {
	fibSequence := []int{X, Y}

	// Generate the Fibonacci sequence up to N terms
	for i := 2; i < N; i++ {
		fibSequence = append(fibSequence, fibSequence[i-1]+fibSequence[i-2])
	}

	sumEven := 0
	sumOdd := 0

	// Calculate the sum of even and odd numbers
	for i := 0; i < N; i++ {
		if fibSequence[i]%2 == 0 {
			sumEven += fibSequence[i]
		} else {
			sumOdd += fibSequence[i]
		}
	}

	return sumEven, sumOdd
}

func main() {
	X := 2
	Y := 3
	N := 3

	sumEven, sumOdd := fibonacciSum(X, Y, N)

	fmt.Println("Sum of even:", sumEven)
	fmt.Println("Sum of odd:", sumOdd)
}
