package main

import "fmt"

func palindrome(input string) bool {
	var palindrome string
	for i := len(input) - 1; i >= 0; i-- {
		palindrome += string(input[i])
	}
	if palindrome == input {
		return true
	}
	return false
}

func main() {
	fmt.Println(palindrome("katak"))
}
