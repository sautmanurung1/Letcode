package main

import (
	"fmt"
)

func calculateProbability(sum int) float64 {
	if sum < 3 || sum > 18 {
		return 0 // Invalid sum
	}

	favorableOutcomes := 0     // Number of favorable outcomes
	totalOutcomes := 6 * 6 * 6 // Total number of outcomes

	// Iterate through all possible combinations of dice rolls
	for dice1 := 1; dice1 <= 6; dice1++ {
		for dice2 := 1; dice2 <= 6; dice2++ {
			for dice3 := 1; dice3 <= 6; dice3++ {
				if dice1+dice2+dice3 == sum {
					favorableOutcomes++
				}
			}
		}
	}

	return float64(favorableOutcomes) / float64(totalOutcomes) // Probability = favorable outcomes / total outcomes
}

func main() {
	sum := 10
	probability := calculateProbability(sum)

	fmt.Printf("The probability of getting a sum of %d is %.4f\n", sum, probability)
}
