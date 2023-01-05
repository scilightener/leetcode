package easy

func containsNearbyDuplicate(nums []int, k int) bool {
	seen := make(map[int]int)
	for idx, val := range nums {
		if value, ok := seen[val]; ok && idx-value <= k {
			return true
		}
		seen[val] = idx
	}
	return false
}
