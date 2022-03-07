package main

import "fmt"

func fizzbuzz(n int32){
	for i := 1; int32(i) < n; i++{
		switch{
		case (i%3 == 0) && (i%5==0):
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}

func main(){
	fizzbuzz(100)
}