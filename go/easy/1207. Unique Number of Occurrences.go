package easy

func uniqueOccurrences(arr []int) bool {
	counter := map[int]int{}
	for _, num := range arr {
		if _, ok := counter[num]; !ok {
			counter[num] = 1
		} else {
			counter[num]++
		}
	}
	set := map[int]struct{}{}
	for _, val := range counter {
		if _, ok := set[val]; ok {
			return false
		}
		set[val] = struct{}{}
	}
	return true
}
