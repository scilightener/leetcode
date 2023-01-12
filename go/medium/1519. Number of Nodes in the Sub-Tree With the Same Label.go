package medium

func countSubTrees(n int, edges [][]int, labels string) []int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}
	ans := make([]int, n)

	var dfs func(int, int) [26]int
	dfs = func(node int, parent int) [26]int {
		total := [26]int{}
		for _, child := range adj[node] {
			if child != parent {
				count := dfs(child, node)
				for i := range count {
					total[i] += count[i]
				}
			}
		}
		total[labels[node]-'a']++
		ans[node] = total[labels[node]-'a']
		return total
	}

	dfs(0, -1)
	return ans
}

/*
[solution is not mine]
complexity:
	time: O(N) for dfs
	space: O(N) for dfs

let's say our dfs func returns an array consisting of the number of labels occurring in the subtree of the given node
ok, now we're done bc all we need is to dfs the root and get the answer!
*/
