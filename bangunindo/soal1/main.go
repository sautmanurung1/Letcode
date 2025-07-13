package main

import (
	"fmt"
	"sort"
)

func minDiff(arr []int32) int32 {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	var minDifference int32 = 0
	for i := 1; i < len(arr); i++ {
		minDifference += abs(arr[i] - arr[i-1])
	}

	return minDifference
}

func abs(x int32) int32 {
	if x < 0 {
		return -x
	}

	return x
}
func main() {
	arr := []int32{5, 1, 3, 7, 3}
	fmt.Println(minDiff(arr)) // Output: 6
}
