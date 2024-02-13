package main

func ways(total int32, k int32) int32 {
	// Write your code here
	const mod int32 = 1000000007
	dp := make([][]int32, total+1)
	for i := range dp {
		dp[i] = make([]int32, k+1)
	}

	for i := int32(0); i <= k; i++ {
		dp[0][i] = 1
	}

	for i := int32(1); i <= total; i++ {
		for j := int32(1); j <= k; j++ {
			dp[i][j] = dp[i][j-1]
			if i >= j {
				dp[i][j] += dp[i-j][j]
			}
		}
	}
	return dp[total][k]
}
