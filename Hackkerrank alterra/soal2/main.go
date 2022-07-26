package main

import (
	"fmt"
)

func getMaximumScore(arr []int32, k int32) int64 {
	// make looping for array
	var arr_len = int32(len(arr))
	var arr_new []int32
	for i := int32(0); i < arr_len; i++ {
		arr_new = append(arr_new, arr[i])
	}
	// check the biggest number in the array
	var max_num int32
	for i := int32(0); i < arr_len; i++ {
		if arr_new[i] > max_num {
			max_num = arr_new[i]
		}
	}
	// sum the array with k to be a length
	var sum_arr int64
	for i := int32(0); i < k; i++ {
		sum_arr += int64(max_num)
	}
	return sum_arr
}

func main() {
	fmt.Println(getMaximumScore([]int32{20, 4, 3, 1, 9}, 4)) // Output : 40
	fmt.Println(getMaximumScore([]int32{4, 5, 18, 1}, 3))    // Output : 29
	fmt.Println(getMaximumScore([]int32{1, 1, 1}, 2))        // Output : 2
}
