package main

func minimumGroups(predators []int32) int32 {
	// Write your code here
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
