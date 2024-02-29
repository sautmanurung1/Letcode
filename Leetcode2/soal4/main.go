package main

import "strings"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}

	var rows []string
	for i := 0; i < numRows && i < len(s); i++ {
		rows = append(rows, "")
	}

	direction := -1 // Direction for zigzag pattern
	row := 0

	for i := 0; i < len(s); i++ {
		rows[row] += string(s[i])
		if row == 0 || row == numRows-1 {
			direction *= -1 // Change direction at the top or bottom row
		}
		row += direction
	}

	var result strings.Builder
	for _, r := range rows {
		result.WriteString(r)
	}

	return result.String()
}
