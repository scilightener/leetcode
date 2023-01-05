package easy

import "sort"

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	pref_sum := make([]int, len(nums))
	pref_sum[0] = nums[0]
	for i, n := range nums[1:] {
		pref_sum[i+1] = pref_sum[i] + n
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		ans[i] = b_search(pref_sum, q) + 1
	}
	return ans
}

func b_search(a []int, u int) int {
	l, r := -1, len(a)
	for r-l > 1 {
		m := (l + r) / 2
		if a[m] <= u {
			l = m
		} else {
			r = m
		}
	}
	return l
}

/*
complexity:
	time: O(NlogN + MlogN) bc we sort nums and iterate trough queries with bsearch on nums
	space: O(N) need pref_sum array

the idea is pretty straight forward: first, sort nums, assuming that smaller numbers you take in the sum, more it can fit in the quaries[i] upper bound (greedy approach)
create prefix sum array to indicate, what quaries[i] at most can be to satisfy proble condition
then just for every query find its place in pref_sum and add 1 since it's 0-indexed array. here is the ans for queries[i]!!
*/
