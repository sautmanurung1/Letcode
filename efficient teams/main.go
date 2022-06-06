package main

func getTotalEfficiency(skill []int32) int64{
	// the sum of all from val slice
	totalEfficiency := int64(0)
	for i := 0; i < len(skill); i++ {
		if totalEfficiency < 0 {
			totalEfficiency = -1 
		}
		if skill[i] < 0 {
			skill[i] = -1
		}
		totalEfficiency += int64(skill[i])
	}
	return totalEfficiency
}

func main(){
	// get total efficiency
	totalEfficiency := getTotalEfficiency([]int32{5, 4, 2, 1}) // 13
	totalEfficiency2 := getTotalEfficiency([]int32{2, 1, 1, 4, 3, 5}) // -1
	// print total efficiency
	println(totalEfficiency)
	println(totalEfficiency2)
}