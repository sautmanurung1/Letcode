package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func jumlahBus(totalKeluarga int, anggotaKeluarga string) string {
	re := regexp.MustCompile(`\s`)
	jumlahKeluarga := re.ReplaceAllString(anggotaKeluarga, "")
	var memberKeluarga []string
	for _, ch := range jumlahKeluarga {
		memberKeluarga = append(memberKeluarga, string(ch))
	}
	if totalKeluarga != len(memberKeluarga) {
		return "Input must be equal with count of family"
	}
	var convertString []int
	for _, str := range memberKeluarga {
		num, _ := strconv.Atoi(str)
		convertString = append(convertString, num)
	}
	totalSemuaOrang := 0
	for _, num := range convertString {
		totalSemuaOrang += num
	}
	if totalSemuaOrang%4 == 0 {
		return fmt.Sprintf("Bus required is: %d", totalSemuaOrang/4)
	} else if totalSemuaOrang%4 != 0 {
		return fmt.Sprintf("Bus required is: %d", (totalSemuaOrang+4-1)/4)
	}
	return ""
}

func main() {
	fmt.Println(jumlahBus(5, "1 2 4 3 3"))
	fmt.Println(jumlahBus(8, "2 3 4 4 2 1 3 1"))
	fmt.Println(jumlahBus(5, "1 5"))
}
