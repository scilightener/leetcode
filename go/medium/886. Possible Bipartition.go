package medium

func possibleBipartition(n int, dislikes [][]int) bool {
	graph, marks := make([][]int, n), make([]int, n)
	for _, p := range dislikes {
		u, v := p[0]-1, p[1]-1
		graph[u] = append(graph[u], v)
		graph[v] = append(graph[v], u)
	}
	for i := 0; i < n; i++ {
		if marks[i] == 0 && !dfs(graph, i, 1, marks) {
			return false
		}
	}
	return true
}

func dfs(graph [][]int, start, mark int, marks []int) bool {
	if marks[start] != 0 {
		return marks[start] == mark
	}
	marks[start] = mark
	for _, v := range graph[start] {
		if !dfs(graph, v, -mark, marks) {
			return false
		}
	}
	return true
}

/*
complexity:
	time: O(N**2) i think but not sure...
	space: O(N**2) if graph is complete

the idea is to mark every node so that it's either in 1 group or in -1 group
dfs func returns whether it's possible to mark current node(start) as
element of the given group(mark).
well, maybe it wasn't the bad idea to name it simply colors...
graph[i] only contains the nodes that should be in different group than i
*/
