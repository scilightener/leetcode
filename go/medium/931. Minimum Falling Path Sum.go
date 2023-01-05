package medium

import "math"

func minFallingPathSum(matrix [][]int) int {
	n := len(matrix)
	dp1, dp := make([]int, n+2), make([]int, n+2)
	dp1[0], dp[0], dp1[n+1], dp[n+1] = 10001, 10001, 10001, 10001
	for _, row := range matrix {
		for j, elem := range row {
			dp[j+1] = elem + min(dp1[j], dp1[j+1], dp1[j+2])
		}
		dp1, dp = dp, dp1
	}
	return min(dp1[1 : n+1]...)
}

func min(items ...int) int {
	minItem := math.MaxInt
	for _, val := range items {
		if val < minItem {
			minItem = val
		}
	}

	return minItem
}
