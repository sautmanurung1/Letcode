package main

import "fmt"

func toChar(i int32) string {
	return string(rune('A' - 1 + i + 33))
}

func slowestKey(keyTimes [][]int32) string {
	// Write your code here
	l := len(keyTimes)
	var interval, key int32
	for i := l - 1; i >= 0; i-- {
		if i > 0 {
			x := keyTimes[i][1] - keyTimes[i-1][1]
			if x > interval {
				interval = x
				key = keyTimes[i][0]
			}
		} else {
			if keyTimes[i][1] > interval {
				interval = keyTimes[i][1]
				key = keyTimes[i][0]
			}
		}
	}
	return toChar(key)
}
func main() {
	a := [][]int32{{0, 2}, {1, 5}, {0, 9}, {2, 15}}
	fmt.Println(slowestKey(a))
}
