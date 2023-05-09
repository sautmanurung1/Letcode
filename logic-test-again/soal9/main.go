package main

import (
	"fmt"
)

func multiplyMatrices(matrix1 [][]int, matrix2 [][]int) [][]int {
	m := len(matrix1)    // Number of rows in matrix1
	n := len(matrix1[0]) // Number of columns in matrix1
	p := len(matrix2[0]) // Number of columns in matrix2

	// Create an empty result matrix
	result := make([][]int, m)
	for i := 0; i < m; i++ {
		result[i] = make([]int, p)
	}

	// Perform matrix multiplication
	for i := 0; i < m; i++ {
		for j := 0; j < p; j++ {
			sum := 0
			for k := 0; k < n; k++ {
				sum += matrix1[i][k] * matrix2[k][j]
			}
			result[i][j] = sum
		}
	}

	return result
}

func main() {
	// Example matrices
	matrix1 := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	matrix2 := [][]int{
		{10, 11},
		{12, 13},
		{14, 15},
	}

	// Multiply the matrices
	resultMatrix := multiplyMatrices(matrix1, matrix2)

	// Print the result
	fmt.Println(resultMatrix)
}
