package main

import "fmt"

func Solution(N int, itemList [][]int, M int, query [][]int) {
	// Your code starts here.
	var count = make([]int, M)
	for i := 0; i < M; i++ {
		for k := 0; k < N; k++ {
			if (query[i][0] <= itemList[k][0] && itemList[k][0] <= query[i][1]) && (query[i][2] <= itemList[k][1] && itemList[k][1] <= query[i][3]) {
				count[i] += 1
			}
		}
	}
	for _, v := range count {
		fmt.Println(v)
	}
}
