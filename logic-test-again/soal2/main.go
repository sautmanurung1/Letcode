package main

import "fmt"

type Food struct {
	name        string
	carbsPer100 int
	pricePer100 int
}

func main() {
	// Define the available foods
	rice := Food{name: "Rice", carbsPer100: 28, pricePer100: 50} // price is in cents
	corn := Food{name: "Corn", carbsPer100: 21, pricePer100: 40}
	potato := Food{name: "Potato", carbsPer100: 17, pricePer100: 30}

	// Define the target number of carbs per day
	targetCarbs := 400

	// Define the budget in cents
	budget := 1000

	// Initialize the variables for the optimal solution
	var minCost int = -1
	var minCombo []int

	// Loop over all possible combinations of foods
	for i := 0; i <= budget/rice.pricePer100; i++ {
		for j := 0; j <= budget/corn.pricePer100; j++ {
			for k := 0; k <= budget/potato.pricePer100; k++ {
				// Calculate the total carbs and cost for this combination
				totalCarbs := i*rice.carbsPer100 + j*corn.carbsPer100 + k*potato.carbsPer100
				totalCost := i*rice.pricePer100 + j*corn.pricePer100 + k*potato.pricePer100

				// Check if this combination meets the carb target and is within budget
				if totalCarbs == targetCarbs && totalCost <= budget {
					// Check if this is the new optimal solution
					if minCost == -1 || totalCost < minCost {
						minCost = totalCost
						minCombo = []int{i, j, k}
					}
				}
			}
		}
	}

	// Print the optimal solution
	fmt.Printf("Optimal food combination:\n")
	fmt.Printf("%s: %d grams\n", rice.name, minCombo[0]*100)
	fmt.Printf("%s: %d grams\n", corn.name, minCombo[1]*100)
	fmt.Printf("%s: %d grams\n", potato.name, minCombo[2]*100)
	fmt.Printf("Total cost: $%.2f\n", float64(minCost)/100.0)
}
