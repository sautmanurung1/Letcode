package main

import "fmt"

func primeX(number int) int {
	var check bool = true
	var prime int = 1
	var counter int = 0
	var i int = 1
	for counter != number {
		i += 1
		if i > 1 {
			for x := 2; x < i; x++ {
				if i%x == 0 {
					check = false
					break
				} else {
					check = true
				}
			}
			if check == true {
				prime = i
				counter += 1
			}

		}

	}
	return prime

}

func main() {
	fmt.Println(primeX(1))
	fmt.Println(primeX(5))
	fmt.Println(primeX(8))
	fmt.Println(primeX(9))
	fmt.Println(primeX(10))
}