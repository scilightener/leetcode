package medium

type Set map[byte]struct{}

func (set Set) Contains(key byte) bool {
	_, ok := set[key]
	return ok
}

func (set Set) Add(key byte) {
	if set.Contains(key) {
		return
	}
	set[key] = struct{}{}
}

func isValidSudoku(board [][]byte) bool {
	usedInRow := make(Set, 0)
	usedInCol := make(Set, 0)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[j][i] != '.' && usedInRow.Contains(board[j][i]) {
				return false
			}
			usedInRow.Add(board[j][i])
			if board[i][j] != '.' && usedInCol.Contains(board[i][j]) {
				return false
			}
			usedInCol.Add(board[i][j])
		}
		usedInRow = make(Set, 0)
		usedInCol = make(Set, 0)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			usedInRow = make(Set, 0)
			for i1 := 0; i1 < 3; i1++ {
				for j1 := 0; j1 < 3; j1++ {
					if board[3*i+i1][3*j+j1] != '.' && usedInRow.Contains(board[3*i+i1][3*j+j1]) {
						return false
					}
					usedInRow.Add(board[3*i+i1][3*j+j1])
				}
			}
		}
	}

	return true
}
