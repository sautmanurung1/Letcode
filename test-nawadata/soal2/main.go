package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Input the number of families: ")
	family, _ := reader.ReadString('\n')
	family = strings.TrimSpace(family)
	fmt.Print("Input the number of members in the family (separated by a space): ")
	member, _ := reader.ReadString('\n')
	member = strings.TrimSpace(member)
	memberSlice := strings.Split(member, " ")
	var memberInt []int
	for _, str := range memberSlice {
		num, _ := strconv.Atoi(str)
		memberInt = append(memberInt, num)
	}
	if len(memberInt) != len(family) {
		fmt.Println("Input must be equal with count of family")
	} else {
		result := 0
		maxMiniBus := 4
		for i, num := range memberInt {
			if num == maxMiniBus && num != 0 {
				result += 1
			} else if num < maxMiniBus && num != 0 {
				sum := num + memberInt[i+1]
				modulus := sum % maxMiniBus
				memberInt[i+1] = 0
				if sum < maxMiniBus {
					memberInt[i+1] = 0
				} else {
					memberInt[i+1] = modulus
				}
				result += 1
			}
		}
		fmt.Printf("Minimum bus required is: %d", result)
	}
}
