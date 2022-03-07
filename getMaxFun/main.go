package main

import "fmt"

func getMaxFun(singer []int32, length []int32) int64 {
	// first make slice singer
	var singerSlice []int32
	for i := 0; i < len(singer); i++ {
		singerSlice = append(singerSlice, singer[i])
	}
	// then make slice length
	var lengthSlice []int32
	for i := 0; i < len(length); i++ {
		lengthSlice = append(lengthSlice, length[i])
	}
	// order the lengthSlice from small to large
	for i := 0; i < len(lengthSlice); i++ {
		for j := i + 1; j < len(lengthSlice); j++ {
			if lengthSlice[i] > lengthSlice[j] {
				lengthSlice[i], lengthSlice[j] = lengthSlice[j], lengthSlice[i]
			}
		}
	}
	// multiply singerslice and lengthSlice to variable max
	var max int64
	for i := 0; i < len(singerSlice); i++ {
			if singerSlice[i] > lengthSlice[i] {
				max += int64(singerSlice[i] * lengthSlice[i])
			} else {
				max += int64(lengthSlice[i] * singerSlice[i])
			}
			fmt.Printf("%d * %d = %d\n", singerSlice[i], lengthSlice[i], max)
	}
	return max
}

func main(){
	println(getMaxFun([]int32{1,1,2}, []int32{5,4,3}))
	println(getMaxFun([]int32{1,2,3}, []int32{1,2,3}))
}