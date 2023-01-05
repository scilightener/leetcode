package medium

func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	maze[entrance[0]][entrance[1]] = '+'

	isExit := func(point [2]int) bool {
		return point[0] == 0 || point[1] == 0 || point[0] == m-1 || point[1] == n-1
	}

	moves := [][]int{{0, 1}, {0, -1}, {-1, 0}, {1, 0}}
	cur, next := [][2]int{}, [][2]int{}
	for _, move := range moves {
		x, y := entrance[0]+move[0], entrance[1]+move[1]
		if x >= 0 && y >= 0 && x < m && y < n && maze[x][y] == '.' {
			cur = append(cur, [2]int{x, y})
			maze[x][y] = '+'
		}
	}
	step := 1
	for len(cur)+len(next) > 0 {
		for _, point := range cur {
			if isExit(point) {
				return step
			}
			for _, move := range moves {
				x, y := point[0]+move[0], point[1]+move[1]
				if x >= 0 && y >= 0 && x < m && y < n && maze[x][y] == '.' {
					next = append(next, [2]int{x, y})
					maze[x][y] = '+'
				}
			}
		}
		cur, next = next, cur[:0]
		step++
	}

	return -1
}
