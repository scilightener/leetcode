package _go

func validPath(n int, edges [][]int, source int, destination int) bool {
	graph := map[int][]int{}
	for _, pair := range edges {
		if _, ok := graph[pair[0]]; ok {
			graph[pair[0]] = append(graph[pair[0]], pair[1])
		} else {
			graph[pair[0]] = []int{pair[1]}
		}
		if _, ok := graph[pair[1]]; ok {
			graph[pair[1]] = append(graph[pair[1]], pair[0])
		} else {
			graph[pair[1]] = []int{pair[0]}
		}
	}

	return dfs(graph, source, destination)
}

func dfs(graph map[int][]int, start, end int) bool {
	visited := map[int]struct{}{}
	cur := []int{start}
	for len(cur) > 0 {
		node := cur[len(cur)-1]
		cur = cur[:len(cur)-1]
		if node == end {
			return true
		}
		if _, ok := visited[node]; ok {
			continue
		}
		visited[node] = struct{}{}
		if adjacent, ok := graph[node]; ok {
			cur = append(cur, adjacent...)
		}
	}
	return false
}

/*
complexity:
	time: O(N) as we iterate through edges to construct graph and then do dfs
	space: O(N**2) in the case if a graph is complete

the main idea is simply dfs.
it might be bfs as well, i'm just too lazy to implement it
it could be union-find... but it's toooooo much more code
and who cares if dfs just works? :)
*/
