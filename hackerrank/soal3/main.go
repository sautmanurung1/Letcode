package main

import "fmt"

func ways(total int, k int) int {
	if k == 1 {
		return 1
	}
	a := make([]int, total)
	currLen := total
	var A [][]int
	x := 1
	i := 0
	for x < k {
		x++
		a[i] = x
		currLen--
		A = append(A, a[:currLen])
		if x == k {
			x = 1
			i++
		}
		if currLen-1 == i || currLen == i {
			break
		}
	}
	return len(A)
}
func main() {
	fmt.Println(ways(5, 3)) // Output : 5
}
