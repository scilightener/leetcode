package medium

import "strconv"

func evalRPN(tokens []string) (num int) {
	nums := []int{}
	ops := map[string]func(int, int) int{
		"+": func(i1, i2 int) int { return i1 + i2 },
		"-": func(i1, i2 int) int { return i1 - i2 },
		"*": func(i1, i2 int) int { return i1 * i2 },
		"/": func(i1, i2 int) int { return i1 / i2 },
	}
	for _, token := range tokens {
		if f, ok := ops[token]; ok {
			num = f(nums[len(nums)-2], nums[len(nums)-1])
			nums = nums[:len(nums)-2]
		} else {
			num, _ = strconv.Atoi(token)
		}
		nums = append(nums, num)
	}
	return
}
