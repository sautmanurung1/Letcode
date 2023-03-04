package main

import "fmt"

func swap(a, b *int) {
	*a, *b = *b, *a
}

func main() {
	a := 50
	b := 63
	swap(&a, &b)
	fmt.Println(a, b)
}
