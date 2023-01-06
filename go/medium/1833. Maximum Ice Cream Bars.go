package medium

import "sort"

func maxIceCream(costs []int, coins int) int {
	sort.Ints(costs)
	for i, cost := range costs {
		if coins < cost {
			return i
		}
		coins -= cost
	}
	return len(costs)
}

/*
complexity:
	time: O(NlogN) for sorting
	space: O(logN) for sorting (depends on language)

actually, don't know why it's medium, i saw easy problems that are way more complicated than this..
just sort costs and take ice creams while their sum is less than coins, not really smth outstanding
*/
