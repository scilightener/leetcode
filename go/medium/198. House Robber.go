package medium

func rob(nums []int) int {
	if len(nums) < 3 {
		return max(nums...)
	}
	dp3, dp2, dp1 := 0, nums[0], max(nums[0], nums[1])
	for i := 2; i < len(nums); i++ {
		dp3, dp2, dp1 = dp2, dp1, nums[i]+max(dp2, dp3)
	}
	return max(dp1, dp2)
}

func max(items ...int) int {
	mx := items[0]
	for _, val := range items {
		if val > mx {
			mx = val
		}
	}
	return mx
}
