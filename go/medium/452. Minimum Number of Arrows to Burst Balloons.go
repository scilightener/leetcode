package _go

import "sort"

func findMinArrowShots(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})
	arrows, right := 1, points[0][1]
	for _, rng := range points[1:] {
		if rng[0] <= right {
			right = min(rng[1], right)
			continue
		}
		arrows++
		right = rng[1]
	}
	return arrows
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

/*
complexity:
	time: O(NlogN) for sorting
	space: O(logN) for sorting (depends on a language)

find the most nested balloon (the one before balloon, whoose left border exceeds our current right border) and then just shoot in it
now, just repeate this operation til there are non-burst balloons
*/
