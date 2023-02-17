package main

import "fmt"

func main() {
	input := []string{"a", "a", "a", "b", "c", "c", "b", "b", "b", "d", "d", "e", "e", "e"}

	obj := make(map[string][]string)
	for _, curr := range input {
		if _, ok := obj[curr]; ok {
			obj[curr] = append(obj[curr], curr)
		} else {
			obj[curr] = []string{curr}
		}
	}

	fmt.Println(obj)
}
