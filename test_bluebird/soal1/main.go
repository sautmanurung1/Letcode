package main

import "fmt"

func countBuildings(arr []int, n int) int {
	count := 1
	curr_max := arr[0]
	for i := 1; i < n; i++ {
		if arr[i] > curr_max || arr[i] == curr_max {
			count += 1
			curr_max = arr[i]
		}
	}
	return count
}

var (
	arr = []int{7, 4, 8, 2, 9}
	n   = len(arr)
)

func main() {
	fmt.Println(countBuildings(arr, n))
}
