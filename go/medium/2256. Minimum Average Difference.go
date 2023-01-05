package medium

import "math"

func minimumAverageDifference(nums []int) int {
	n, s := len(nums), sum(nums)
	frwd_sum := make([]int, n)
	frwd_sum[0] = nums[0]
	for i := 1; i < n; i++ {
		frwd_sum[i] = frwd_sum[i-1] + nums[i]
	}

	mn := math.MaxInt
	minIdx := -1
	for i := 0; i < n-1; i++ {
		h := abs(frwd_sum[i]/(i+1) - (s-frwd_sum[i])/(n-i-1))
		if h < mn {
			minIdx = i
			mn = h
		}
	}

	if s/n < mn {
		return len(nums) - 1
	}

	return minIdx
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sum(a []int) int {
	s := 0
	for _, val := range a {
		s += val
	}

	return s
}
