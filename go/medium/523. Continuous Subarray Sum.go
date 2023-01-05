package medium

func checkSubarraySum(nums []int, k int) bool {
	if len(nums) < 2 {
		return false
	}
	if len(nums) == 2 {
		if (nums[0]+nums[1])%k == 0 {
			return true
		}
		return false
	}

	rems := make(map[int]struct{})
	rems[0] = struct{}{}
	currem, cursum := 0, nums[0]+nums[1]
	if cursum%k == 0 {
		return true
	}
	for i := 2; i < len(nums); i++ {
		currem += nums[i-2]
		rems[currem%k] = struct{}{}
		cursum += nums[i]
		if _, ok := rems[cursum%k]; ok {
			return true
		}
	}
	return false
}
