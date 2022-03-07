package main

import "fmt"

func moneyCoins(money int) []int{
	coins := []int{10000, 5000, 2000, 500, 200, 100, 50, 20, 10, 5, 1}
	var result []int
	for _, coin := range coins {
		for money >= coin {
			money -= coin
			result = append(result, coin)
		}
	}
	return result
}

func main(){
	fmt.Println(moneyCoins(123)) // [100, 20, 1, 1, 1]
	fmt.Println(moneyCoins(432)) // [200, 200, 20, 10, 1, 1]
	fmt.Println(moneyCoins(543)) // [500, 20, 20, 1, 1, 1]
	fmt.Println(moneyCoins(7752)) // [5000, 2000, 500, 200, 50, 1, 1]
	fmt.Println(moneyCoins(15321)) // [10000, 5000, 200, 100 , 20, 1]
}