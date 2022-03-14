package main

func reverseArray(arr []int32) []int32{
	var newArr []int32
	for i := len(arr)-1; i >= 0; i--{
		newArr = append(newArr, arr[i])
	}
	return newArr
}