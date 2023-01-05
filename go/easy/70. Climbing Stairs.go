package easy

func climbStairs(n int) int {
	if n < 3 {
		return n
	}
	dp1, dp2 := 2, 1
	for i := 2; i < n; i++ {
		dp1, dp2 = dp1+dp2, dp1
	}
	return dp1
}
