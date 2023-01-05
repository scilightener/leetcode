package hard

type State struct {
	i int
	j int
	k int
}

var moves = [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
var cur = make([]State, 0, 1)
var next = make([]State, 0, 1)

func shortestPath(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	if m == 1 && n == 1 {
		return 0
	}

	visited := [40][40]int{}
	visited[0][0] = k + 1

	cur = cur[:1]
	cur[0] = State{0, 0, k}

	next = next[:0]

	for result := 1; len(cur) != 0; result++ {
		for _, position := range cur {
			for _, move := range moves {
				next_i := position.i + move[0]
				if next_i < 0 || next_i >= m {
					continue
				}

				next_j := position.j + move[1]
				if next_j < 0 || next_j >= n {
					continue
				}

				if next_i == m-1 && next_j == n-1 {
					return result
				}

				next_k := position.k
				if grid[next_i][next_j] == 1 {
					if next_k == 0 {
						continue
					}
					next_k--
				}

				if visited[next_i][next_j] <= next_k {
					visited[next_i][next_j] = next_k + 1
					next = append(next, State{next_i, next_j, next_k})
				}
			}
		}
		cur, next = next, cur[:0]
	}

	return -1
}
