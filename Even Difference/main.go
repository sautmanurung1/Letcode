package main

func findLongestSubsequence(arr []int32) int32{
	var max int32
	var current int32
	for i := 0; i < len(arr); i++{
		current = 1
		for j := i+1; j < len(arr); j++{
			if arr[i] == arr[j]{
				current++
			}
		}
		if current > max{
			max = current
		}
	}
	return max
}