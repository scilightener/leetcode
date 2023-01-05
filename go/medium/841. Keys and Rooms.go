package medium

func canVisitAllRooms(rooms [][]int) bool {
	graph := map[int][]int{}
	for room, keys := range rooms {
		graph[room] = keys
	}
	return dfs(graph, 0, len(rooms))
}

func dfs(graph map[int][]int, start, n int) bool {
	visited := map[int]struct{}{}
	cur := []int{start}
	for len(cur) > 0 {
		node := cur[len(cur)-1]
		cur = cur[:len(cur)-1]
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}
		cur = append(cur, graph[node]...)
		if len(visited) == n {
			return true
		}
	}

	return false
}

/*
complexity:
	time: O(N) simple dfs
	space: O(N**2) for graph map (if graph is complete)

the idea is still just dfs. could be bfs as well
*/
