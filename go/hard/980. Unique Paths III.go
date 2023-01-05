package hard

var (
	moves   []int = []int{0, -1, 0, 1, 0}
	grid    [][]int
	targetX int
	targetY int
)

func uniquePathsIII(_grid [][]int) int {
	cells, x, y := 2, 0, 0
	grid = _grid
	for i, row := range grid {
		for j := range row {
			if grid[i][j] == 0 {
				cells++
			} else if grid[i][j] == 1 {
				x, y = i, j
			} else if grid[i][j] == 2 {
				targetX, targetY = i, j
			}
		}
	}
	return backtrack(x, y, cells)
}

func backtrack(r, c, cnt int) (ans int) {
	cnt--
	if cnt < 0 {
		return 0
	}
	if cnt == 0 && r == targetX && c == targetY {
		return 1
	}
	grid[r][c] = -1
	for i := range moves[:4] {
		x, y := r+moves[i], c+moves[i+1]
		if 0 <= x && x < len(grid) &&
			0 <= y && y < len(grid[0]) &&
			grid[x][y] != -1 {
			ans += backtrack(x, y, cnt)
		}
	}
	grid[r][c] = 0
	return ans
}

/*
complexity:
	time: O(??) not sure but maybe O(3**MN)? choose three directions to go every step or something (skip the direction we came from)... but it's very rough estimate though
	space: O(??) dunno actually sorry

let's say our backtrack function returns the number of ways to go from (r, c) to (targetX, targetY) with each way length = cnt

then we just enumerate through all the possible directions and call backtrack of new start position with cnt-1
don't forget about the base case when working with recursion!!
to prevent going one cell twice, mark the current cell as the obstacles by grid[r][c] = -1 and mark them back when we're done
easy, not hard!
*/
