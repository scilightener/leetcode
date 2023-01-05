package medium

func exist(board [][]byte, word string) bool {
	m, n := len(board), len(board[0])
	var dfs func(r, c, idx int) bool
	dfs = func(r, c, idx int) bool {
		if idx == len(word) {
			return true
		}
		if r < 0 || r >= m || c < 0 || c >= n ||
			board[r][c] != word[idx] || board[r][c] == '*' {
			return false
		}

		visited := board[r][c]
		board[r][c] = '*'
		r0, r1 := dfs(r-1, c, idx+1), dfs(r+1, c, idx+1)
		c0, c1 := dfs(r, c-1, idx+1), dfs(r, c+1, idx+1)
		board[r][c] = visited
		return r0 || r1 || c0 || c1
	}

	for rowIdx := range board {
		for colIdx := range board[0] {
			if dfs(rowIdx, colIdx, 0) {
				return true
			}
		}
	}

	return false
}
