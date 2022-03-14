package main

import "fmt"

func main(){
	var kecepatan int = 2 // m/detik

	var jarak int = 0

	var waktu int = (3600 * 2) + (60 * 10) + 17

	jarak = kecepatan * waktu

	fmt.Println(jarak, "Meter")
}