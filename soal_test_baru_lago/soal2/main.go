package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func calculateMinMaxSum(arr []int) string {
	// Sort the array in ascending order
	sort.Ints(arr)
	// Calculate the sum of the four smallest numbers
	minSum := arr[0] + arr[1] + arr[2] + arr[3]
	// Calculate the sum of the four largest numbers
	maxSum := arr[1] + arr[2] + arr[3] + arr[4]
	return strconv.Itoa(minSum) + " " + strconv.Itoa(maxSum)
}

func main() {
	// Read input from the user
	var input string
	fmt.Scanln(&input)
	// Split the input string into an array of integers
	arrStr := strings.Split(input, " ")
	arr := make([]int, len(arrStr))
	for i, str := range arrStr {
		arr[i], _ = strconv.Atoi(str)
	}
	// Calculate the minimum and maximum sums
	result := calculateMinMaxSum(arr)
	// Print the result
	fmt.Println(result)
}
