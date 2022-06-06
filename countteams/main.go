package main

func countTeams(rating []int32, queries [][]int32) []int32 {
    // Write your code here
	var result []int32
	for i := 0; i < len(queries); i++ {
		var count int32
		for j := 0; j < len(rating); j++ {
			for k := j + 1; k < len(rating); k++ {
				if rating[j] > rating[k] && queries[i][0] > rating[j] && queries[i][1] > rating[k] {
					count++
				}
				if rating[j] < rating[k] && queries[i][0] < rating[j] && queries[i][1] < rating[k] {
					count++
				}
			}
		}
		result = append(result, count)
	}
	return result
}

func main(){
}