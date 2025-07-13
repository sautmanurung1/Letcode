package main

func minIndex(steps int32, badIndex int32) int32 {
	// Set initial position to 0
	position := int32(0)

	// Loop through steps
	for i := int32(0); i < steps; i++ {
		// If the current position is the bad index, skip it
		if position == badIndex {
			position++
		}
		// Move to the next position
		position += i + 1
	}

	// Return the maximum index that can be reached
	return position
}
