package easy

import "sort"

func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[j] < nums[i]
	})

	ans := 0
	for i := 0; i < len(nums)-2; i++ {
		if p := perimeter(nums[i : i+3]); p > ans {
			return p
		}
	}
	return 0
}

func perimeter(a []int) int {
	if s := a[0] + a[1] + a[2]; s-a[0] > a[0] {
		return s
	}
	return 0
}
