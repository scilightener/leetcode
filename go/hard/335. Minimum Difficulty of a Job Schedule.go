package hard

type State struct {
	i int
	d int
}

var days int

func minDifficulty(ds []int, d int) int {
	if len(ds) < d {
		return -1
	}
	mem := make(map[State]int)
	days = d
	return solve(0, 1, ds, mem)
}

func solve(i int, d int, ds []int, mem map[State]int) int {
	state := State{i, d}
	if res, ok := mem[state]; ok {
		return res
	}

	if d == days {
		res := maxSlice(ds[i:])
		mem[state] = res
		return res
	}

	cur, res := 0, 10000
	for j := i; j < len(ds)-days+d; j++ {
		cur = maxInt(cur, ds[j])
		res = minInt(res, cur+solve(j+1, d+1, ds, mem))
	}

	mem[state] = res
	return res
}

func maxSlice(s []int) int {
	mx := 0
	for _, val := range s {
		mx = maxInt(mx, val)
	}
	return mx
}

func minInt(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}
