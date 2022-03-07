package main

import "fmt"

func BinarySearch(array []int, x int){
	low := 0
	high := len(array) - 1
	for low <= high {
		mid := (low + high) / 2
		if array[mid] == x {
			fmt.Println(mid)
			return
		}
		if array[mid] > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	fmt.Println(-1)
}

func main(){
	BinarySearch([]int{1, 1, 3, 5, 5, 6, 7}, 3) // Output : 2
	BinarySearch([]int{1, 2, 3, 5, 6, 8, 10}, 5) // Output : 3
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 53) // Output : 6
	BinarySearch([]int{12, 15, 15, 19, 24, 31, 53, 59, 60}, 100) // Output : -1
}