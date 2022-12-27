package _go

import "sort"

func maximumBags(capacity []int, rocks []int, additionalRocks int) int {
	for i, rock := range rocks {
		capacity[i] -= rock
	}
	sort.Ints(capacity)
	ans := 0
	for _, val := range capacity {
		if additionalRocks-val >= 0 {
			ans++
			additionalRocks -= val
		} else {
			break
		}
	}
	return ans
}

/*
complexity:
	time: O(NlogN) for sorting
	space: O(1) or O(N) for sorting

the idea is pretty greedy. we first calculate what amount of rocks is left for every bag (capacity[i] - rocks[i]) and then just sort them from the less full to the most. we start fulfilling them from the start. when additionalRocks have ended, return ans
*/
