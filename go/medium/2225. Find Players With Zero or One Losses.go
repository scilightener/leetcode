package medium

import "sort"

type Set map[int]struct{}

func (set *Set) Contains(key int) bool {
	_, ok := (*set)[key]
	return ok
}

func (set *Set) Remove(key int) {
	if !(*set).Contains(key) {
		return
	}
	delete(*set, key)
}

func (set *Set) Add(key int) {
	if (*set).Contains(key) {
		return
	}
	(*set)[key] = struct{}{}
}

func (set *Set) ToIntSlice() []int {
	keys := make([]int, 0, len(*set))

	for k := range *set {
		keys = append(keys, k)
	}

	return keys
}

func findWinners(matches [][]int) [][]int {
	lost0 := make(Set, 0)
	lost1 := make(Set, 0)
	lostMany := make(Set, 0)
	for _, pair := range matches {
		winner, loser := pair[0], pair[1]
		if !lostMany.Contains(winner) && !lost1.Contains(winner) {
			lost0.Add(winner)
		}
		if lost0.Contains(loser) {
			lost0.Remove(loser)
			lost1.Add(loser)
		} else if lost1.Contains(loser) {
			lost1.Remove(loser)
			lostMany.Add(loser)
		} else if lostMany.Contains(loser) {
			continue
		} else {
			lost1.Add(loser)
		}
	}
	ans0, ans1 := lost0.ToIntSlice(), lost1.ToIntSlice()
	sort.Ints(ans0)
	sort.Ints(ans1)
	return [][]int{ans0, ans1}
}
