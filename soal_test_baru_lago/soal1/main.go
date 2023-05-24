package main

import (
	"fmt"
	"sort"
)

func miniMaxSum(arr []int32) {
	// Sort the array in ascending order
	sort.Slice(arr, func(i, j int) bool { return arr[i] < arr[j] })

	// Calculate the minimum sum by summing the first four elements
	var minSum int64 = 0
	for i := 0; i < 4; i++ {
		minSum += int64(arr[i])
	}

	// Calculate the maximum sum by summing the last four elements
	var maxSum int64 = 0
	for i := len(arr) - 1; i > len(arr)-5; i-- {
		maxSum += int64(arr[i])
	}

	// Print the minimum and maximum sums as a single line
	fmt.Println(minSum, maxSum)
}

func main() {
	arr := []int32{1, 3, 5, 7, 9}
	miniMaxSum(arr)
}
