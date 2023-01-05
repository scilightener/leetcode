package hard

import "sort"

type Set map[[2]int]struct{}

func (set Set) Contains(elem [2]int) bool {
	_, in := set[elem]
	return in
}

func (set *Set) Add(elem [2]int) {
	if set.Contains(elem) {
		return
	}
	(*set)[elem] = struct{}{}
}

func (set Set) ToSlice() [][]int {
	keys := [][]int{}
	for k := range set {
		keys = append(keys, []int{k[0], k[1]})
	}

	return keys
}

func outerTrees(trees [][]int) [][]int {
	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0] == trees[j][0] {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})

	set := make(Set, 0)
	for _, point := range findShell(trees) {
		set.Add([2]int{point[0], point[1]})
	}

	return set.ToSlice()
}

func findShell(points [][]int) [][]int {
	top, bottom := [][]int{}, [][]int{}
	for _, point := range points {
		for len(top) > 1 && crossProduct(top[len(top)-2], top[len(top)-1], point) > 0 {
			top = top[:len(top)-1]
		}
		for len(bottom) > 1 && crossProduct(bottom[len(bottom)-2], bottom[len(bottom)-1], point) < 0 {
			bottom = bottom[:len(bottom)-1]
		}
		top, bottom = append(top, point), append(bottom, point)
	}

	return append(top, bottom...)
}

func crossProduct(p0, p1, p2 []int) int {
	return (p1[0]-p0[0])*(p2[1]-p0[1]) - (p1[1]-p0[1])*(p2[0]-p0[0])
}
