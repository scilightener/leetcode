package medium

func maxDistance(grid [][]int) int {
	n := len(grid)

	var firstOne func(int, int) int
	firstOne = func(x, y int) int {
		for k := 1; k <= 2*(n-1); k++ {
			for i := 0; i < n; i++ {
				abs_j := k - abs(x-i)
				if abs_j < 0 {
					continue
				}
				if abs_j+y < n {
					if grid[i][y+abs_j] == 1 {
						return k - 1
					}
				}
				if y-abs_j >= 0 {
					if grid[i][y-abs_j] == 1 {
						return k - 1
					}
				}
			}
		}
		return 2 * (n - 1)
	}

	mx := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				mx = max(mx, firstOne(i, j))
			}
		}
	}

	if mx == 0 {
		mx = -1
	}
	return mx
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
