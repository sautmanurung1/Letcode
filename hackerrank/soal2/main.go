package main

import "fmt"

func minimumGroups(predators []int32) int32 {
	var totalGroup int32 = 1
	num := int32(len(predators))
	for i := int32(0); i < num; i++ {
		species := i
		tempGroup := int32(1)
		for predators[species] > -1 && predators[species] < num && tempGroup < num {
			species = predators[species]
			tempGroup++
		}

		if tempGroup > totalGroup {
			totalGroup = tempGroup
		}
	}

	return totalGroup
}
func main() {
	var predators []int32 = []int32{-1, 8, 6, 0, 7, 3, 8, 9, -1, 6, 1}
	fmt.Println(minimumGroups(predators))
}
