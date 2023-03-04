package main

// import package fmt untuk bisa menampilkan output ke layar
import "fmt"

func main() {
	// loop untuk menampilkan bilangan 1 sampai 100
	for i := 1; i <= 100; i++ {
		isPrime := true // inisialisasi variabel isPrime dengan true
		// loop untuk mengecek apakah bilangan tersebut prima atau tidak
		for j := 2; j < i; j++ {
			if i%j == 0 {
				isPrime = false // jika bilangan habis dibagi j, berarti bukan bilangan prima
				break
			}
		}
		// jika bilangan tersebut prima, tampilkan "Bilangan prima: " diikuti dengan bilangan tersebut
		if isPrime && i != 1 {
			fmt.Println("Bilangan prima:", i)
		}
		// jika bilangan tersebut merupakan kelipatan 9, tampilkan "Kelipatan 9 ke- " diikuti dengan kelipatan 9 tersebut
		if i%9 == 0 {
			fmt.Println("Kelipatan 9 ke-", i/9)
		}
		// jika bukan bilangan prima atau kelipatan 9, tampilkan bilangan tersebut saja
		if !isPrime && i%9 != 0 {
			fmt.Println(i)
		}
	}
}
