package _go

func canJump(nums []int) bool {
	disst := nums[0]
	for _, jump := range nums[1:] {
		if disst <= 0 {
			return false
		}
		disst = max(disst-1, jump)
	}
	return true
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
complexity:
	time: O(N) simple for loop
	space: O(1) just one variable

idea is simple: at the very iteration we know how far we can jump right now
every step this number decreases by one or updates to nums[i] if we have ability to jump to it (disst > 0)
*/
