package main

import "sort"


func lastStoneWeight(weights []int32) int32 {
	// make slice weights
	// sort slice weights
	// while len(weights) > 1
	// 	a := weights[len(weights)-1]
	// 	b := weights[len(weights)-2]
	// 	if a == b {
	// 		weights = weights[:len(weights)-2]
	// 	} else {
	// 		weights = weights[:len(weights)-2]
	// 		weights = append(weights, a-b)
	// 		sort.Ints(weights)
	// 	}
	// return weights[0]
	for len(weights) > 1 {
		a := weights[len(weights)-1]
		b := weights[len(weights)-2]
		if a == b {
			weights = weights[:len(weights)-2]
		} else {
			weights = weights[:len(weights)-2]
			weights = append(weights, a-b)
			sort.Ints(weights)
		}
	}
	return weights[0]
}

func main(){
	println(lastStoneWeight([]int32{2,7,4,1,8,1}))
	println(lastStoneWeight([]int32{2, 4, 5}))
}