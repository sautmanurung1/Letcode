package main

func getMaximumEvenSum(val []int32) int64 {
	// the sum of all from val slice
	totalEfficiency := int64(0)
	for i := 0; i < len(val); i++ {
		if totalEfficiency < 0 {
			totalEfficiency = -1 
		}
		if val[i] < 0 {
			val[i] = -1
		}
		totalEfficiency += int64(val[i])
	}
	return totalEfficiency
}

func main(){
	// get total efficiency
	getmaximum := getMaximumEvenSum([]int32{6, 3, 4, -1, 9, 17}) // 38
	getmaximum2 := getMaximumEvenSum([]int32{-1, -2, -3, 8, 7}) // 14
	getmaximum3 := getMaximumEvenSum([]int32{2, 3, 6, -5, 10, 1, 1}) // 22
	// print total efficiency
	println(getmaximum)
	println(getmaximum2)
	println(getmaximum3)
}