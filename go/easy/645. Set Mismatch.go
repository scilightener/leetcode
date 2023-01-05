package easy

func findErrorNums(nums []int) []int {
	set := make(map[int]struct{})
	res := make([]int, 2)
	s := sum(nums)
	for _, val := range nums {
		if _, ok := set[val]; ok {
			res[0] = val
			break
		}
		set[val] = struct{}{}
	}
	res[1] = (len(nums)+1)*len(nums)/2 - s + res[0]
	return res
}

func sum(nums []int) int {
	s := 0
	for _, val := range nums {
		s += val
	}
	return s
}
