package _go

func dailyTemperatures(temperatures []int) []int {
	mono_stack := make([][2]int, 0)
	ans := make([]int, len(temperatures))
	for i, temp := range temperatures {
		for len(mono_stack) > 0 && mono_stack[len(mono_stack)-1][0] < temp {
			last := mono_stack[len(mono_stack)-1]
			ans[last[1]] = i - last[1]
			mono_stack = mono_stack[:len(mono_stack)-1]
		}
		mono_stack = append(mono_stack, [2]int{temp, i})
	}
	return ans
}
