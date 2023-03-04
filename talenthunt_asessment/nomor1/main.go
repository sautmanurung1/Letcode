package main

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

func BitmapHoles(strArr []string) string {

	// Variables to keep track of visited nodes, and holes
	visited := make(map[int]bool)
	holes := 0
	result := ""

	// Iterate through the 2D matrix
	for i := 0; i < len(strArr); i++ {
		for j := 0; j < len(strArr[i]); j++ {

			// If the current node is 0 and not visited
			if strArr[i][j] == '0' && !visited[i*len(strArr[i])+j] {

				// Do a BFS to find the connected component
				queue := list.New()
				queue.PushBack([2]int{i, j})

				// Mark all the nodes in the same component as visited
				for queue.Len() > 0 {
					coords := queue.Front().Value.([2]int)
					x := coords[0]
					y := coords[1]
					visited[x*len(strArr[x])+y] = true
					queue.Remove(queue.Front())

					// Check the top
					if x > 0 && strArr[x-1][y] == '0' && !visited[(x-1)*len(strArr[x-1])+y] {
						queue.PushBack([2]int{x - 1, y})
					}

					// Check the bottom
					if x < len(strArr)-1 && strArr[x+1][y] == '0' && !visited[(x+1)*len(strArr[x+1])+y] {
						queue.PushBack([2]int{x + 1, y})
					}

					// Check the left
					if y > 0 && strArr[x][y-1] == '0' && !visited[x*len(strArr[x])+y-1] {
						queue.PushBack([2]int{x, y - 1})
					}

					// Check the right
					if y < len(strArr[x])-1 && strArr[x][y+1] == '0' && !visited[x*len(strArr[x])+y+1] {
						queue.PushBack([2]int{x, y + 1})
					}
				}

				// Increment the number of holes
				holes += 1
				result += strconv.Itoa(holes) + ","
			}
		}
	}

	return result
}

func SumHoles(strArr []string) int {
	// Call the BitmapHoles function to get the result string
	result := BitmapHoles(strArr)

	// Split the result string into an array of strings
	arrResult := strings.Split(result, ",")

	// Initialize the sum variable
	sum := 0

	// Iterate through the array of strings
	for _, numString := range arrResult {
		// Parse the string as an integer
		numInt, err := strconv.Atoi(numString)

		// If there is no error, add the number to the sum
		if err == nil {
			sum += numInt
		}
	}

	// Return the sum
	return sum
}

func main() {
	fmt.Println(SumHoles([]string{"1011", "0010"}))
	fmt.Println(SumHoles([]string{"01111", "01101", "00011", "11110"}))
}
