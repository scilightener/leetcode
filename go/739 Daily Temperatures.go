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

/*
complexity:
	time: O(N) as every element happens to be in the stack once and never again
	space: O(N) for the stack

the main idea is to use decreasing monotonic stack:
we store element's value as well as it's index in the stack.
when the new element is proceeding,
we pop all the elements ontop the stack that are less than current
and immediately save answer for these popped (cur_index - their_index)
why is it working? well, because we know that before current element
there were no elements, that were greater than those in the stack,
because otherwise, our alorithm'd kick them off at the first occurance.
*/
