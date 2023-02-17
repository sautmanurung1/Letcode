package main

import "fmt"

func main() {
	arr := []string{"a", "a", "a", "b", "c", "c", "b", "b", "b", "d", "d", "e", "e", "e"}
	count := make(map[string]int)
	for _, i := range arr {
		count[i] = count[i] + 1
	}
	fmt.Println(count)
}
