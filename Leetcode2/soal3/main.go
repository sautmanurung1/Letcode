package main

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func helper(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	return right - left - 1
}

func longestPalindrome(s string) string {
	left := -1
	right := -2

	for i := 0; i < len(s); i++ {
		maxString := max(helper(s, i, i), helper(s, i, i+1))
		if maxString > right-left+1 {
			right = i + maxString/2
			left = right - maxString + 1
		}
	}

	if left >= 0 {
		return s[left : right+1]
	}
	return ""
}
