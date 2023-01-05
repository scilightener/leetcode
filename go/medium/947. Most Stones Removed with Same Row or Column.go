package medium

func removeStones(stones [][]int) int {
	rows, cols := map[int][]int{}, map[int][]int{}
	visited := make([]bool, len(stones))
	for i, point := range stones {
		rows[point[0]] = append(rows[point[0]], i)
	}

	for i, point := range stones {
		cols[point[1]] = append(cols[point[1]], i)
	}

	stonesLeft := 0
	for i := range stones {
		if visited[i] {
			continue
		}
		visited[i] = true
		stonesLeft++
		dfs(rows, cols, stones, visited, i)
	}

	return len(stones) - stonesLeft
}

func dfs(rows, cols map[int][]int, stones [][]int, visited []bool, idx int) {
	visited[idx] = true
	for _, i := range rows[stones[idx][0]] {
		if visited[i] {
			continue
		}
		dfs(rows, cols, stones, visited, i)
	}

	for _, i := range cols[stones[idx][1]] {
		if visited[i] {
			continue
		}
		dfs(rows, cols, stones, visited, i)
	}
}
