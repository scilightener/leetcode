package medium

func findBall(grid [][]int) []int {
	m, n := len(grid), len(grid[0])
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = i
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dp[j] == -1 {
				continue
			}
			new_j := dp[j] + grid[i][dp[j]]
			if new_j >= n {
				new_j = -1
			}
			if new_j == -1 || grid[i][new_j]+new_j == dp[j] {
				dp[j] = -1
				continue
			}

			dp[j] = new_j
		}
	}

	return dp
}
