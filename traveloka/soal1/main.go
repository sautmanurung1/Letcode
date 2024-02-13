package main

func slowestKey(keyTimes [][]int32) string {
	// Write your code here
	var maxDuration int32
	var slowestKey int32

	for i := 1; i < len(keyTimes); i++ {
		duration := keyTimes[i][1] - keyTimes[i-1][1]
		if duration > maxDuration {
			maxDuration = duration
			slowestKey = keyTimes[i][0]
		}
	}

	return string('a' + slowestKey)
}
