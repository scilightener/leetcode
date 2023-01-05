package hard

func numberOfArithmeticSlices(nums []int) int {
	dp := make([]map[int]int, len(nums))
	res := 0
	for i, a := range nums {
		dp[i] = map[int]int{}
		for j, b := range nums[:i] {
			dp[i][a-b] += dp[j][a-b] + 1
			res += dp[j][a-b]
		}
	}

	return res
}
