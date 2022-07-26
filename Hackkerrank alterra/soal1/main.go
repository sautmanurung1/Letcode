package main

import "fmt"

func findLongestSubsequence(arr []int32) int32 {
	// make looping for array
	var arr_len = int32(len(arr))
	var arr_new []int32
	for i := int32(0); i < arr_len; i++ {
		arr_new = append(arr_new, arr[i])
	}
	return arr_len - 1
}

func main() {
	fmt.Println(findLongestSubsequence([]int32{7, 5, 6, 2, 3, 2, 4})) // Output : 6
}
