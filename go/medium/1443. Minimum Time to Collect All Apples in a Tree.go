package medium

func minTime(n int, edges [][]int, hasApple []bool) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	visited := make([]bool, n)
	var dfs func(int) int
	dfs = func(u int) (cnt int) {
		visited[u] = true
		for _, v := range adj[u] {
			if !visited[v] {
				cnt += dfs(v)
			}
		}
		if (hasApple[u] || cnt > 0) && u != 0 {
			cnt += 1
		}
		return cnt
	}
	return dfs(0) * 2
}

/*
[solution is not mine]
complexity:
	time: O(N) dfs
	space: O(N) dfs

first, build the graph (adj)
then, run simple dfs to find the path from 0 to every apple
then multiply by 2 and get the answer
*/
