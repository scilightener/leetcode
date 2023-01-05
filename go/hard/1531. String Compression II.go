package hard

import "math"

type State struct {
	i int
	j int
	k int
	c int
}

func getLengthOfOptimalCompression(s string, k int) int {
	mem := make(map[State]int)
	return solve(0, -1, k, 0, &s, &mem)
}

func solve(i int, j int, k int, c int, s *string, mem *map[State]int) int {
	if k < 0 {
		return math.MaxInt
	}
	if i == len(*s) {
		return 0
	}
	state := State{i, j, k, c}
	if res, ok := (*mem)[state]; ok {
		return res
	}
	if j >= 0 && (*s)[i] == (*s)[j] {
		var cIsOver int
		if c == 1 || c == 9 || c == 99 {
			cIsOver = 1
		}
		res := cIsOver + solve(i+1, i, k, c+1, s, mem)
		(*mem)[state] = res
		return res
	}
	res := min(1+solve(i+1, i, k, 1, s, mem), solve(i+1, j, k-1, c, s, mem))
	(*mem)[state] = res
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
