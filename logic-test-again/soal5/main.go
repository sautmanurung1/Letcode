package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	boxes := []string{"A", "B", "C", "D", "E"}

	var targetBox string
	fmt.Print("Enter the target box (A, B, C, D, or E): ")
	fmt.Scan(&targetBox)

	result := findObject(targetBox, boxes)
	fmt.Println(result)
}

func findObject(targetBox string, boxes []string) string {
	targetIndex := -1
	for i, box := range boxes {
		if box == targetBox {
			targetIndex = i
			break
		}
	}

	if targetIndex == -1 {
		return "Invalid target box."
	}

	rand.Seed(time.Now().UnixNano())
	pointer := rand.Intn(len(boxes))
	steps := []string{boxes[pointer]}

	for i := 0; i < 7; i++ {
		adjacentBoxes := getAdjacentBoxes(pointer, len(boxes))
		pointer = adjacentBoxes[rand.Intn(len(adjacentBoxes))]
		steps = append(steps, boxes[pointer])

		if pointer == targetIndex {
			return fmt.Sprintf("Object found in box %s in %d steps. Steps taken: %s", boxes[pointer], i+1, joinSteps(steps))
		}
	}

	return "Object not found within 7 steps."
}

func getAdjacentBoxes(pointer int, numBoxes int) []int {
	adjacentBoxes := []int{}

	if pointer > 0 {
		adjacentBoxes = append(adjacentBoxes, pointer-1)
	}

	if pointer < numBoxes-1 {
		adjacentBoxes = append(adjacentBoxes, pointer+1)
	}

	return adjacentBoxes
}

func joinSteps(steps []string) string {
	return fmt.Sprintf("%s", steps)
}
