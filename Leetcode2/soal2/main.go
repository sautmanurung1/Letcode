package main

import "math"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	m, n := len(nums1), len(nums2)
	low, high := 0, m

	for low <= high {
		partitionX := (low + high) / 2
		partitionY := (m+n+1)/2 - partitionX

		maxX := math.MinInt64
		if partitionX > 0 {
			maxX = nums1[partitionX-1]
		}

		minX := math.MaxInt64
		if partitionX < m {
			minX = nums1[partitionX]
		}

		maxY := math.MinInt64
		if partitionY > 0 {
			maxY = nums2[partitionY-1]
		}

		minY := math.MaxInt64
		if partitionY < n {
			minY = nums2[partitionY]
		}

		if maxX <= minY && maxY <= minX {
			if (m+n)%2 == 0 {
				return (float64(max(maxX, maxY)) + float64(min(minX, minY))) / 2.0
			}
			return float64(max(maxX, maxY))
		} else if maxX > minY {
			high = partitionX - 1
		} else {
			low = partitionX + 1
		}
	}

	return 0.0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
