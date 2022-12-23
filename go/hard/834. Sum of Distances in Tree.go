package _go

var (
	graph         map[int][]int
	childrencount []int
	ans           []int
)

func sumOfDistancesInTree(n int, edges [][]int) []int {
	graph = make(map[int][]int, n)
	childrencount = make([]int, n)
	ans = make([]int, n)

	for _, p := range edges {
		u, v := p[0], p[1]
		graph[v] = append(graph[v], u)
		graph[u] = append(graph[u], v)
	}

	getChildrenCount(0, -1)
	getAns(0, -1)
	return ans
}

func getChildrenCount(root, parent int) {
	for _, v := range graph[root] {
		if v == parent {
			continue
		}
		getChildrenCount(v, root)
		ans[root] += ans[v] + childrencount[v]
		childrencount[root] += childrencount[v]
	}
	childrencount[root]++
}

func getAns(root, parent int) {
	n := len(ans)
	for _, v := range graph[root] {
		if v == parent {
			continue
		}
		ans[v] = ans[root] + n - 2*(childrencount[v])
		getAns(v, root)
	}
}

/*
complexity:
    time: O(N) bc just 2 dfs
    space: O(N) bc we store just the tree and additional info about every node

ughh i'm not sure i can explain the main idea in several lines...
well, first we calculate the number of children of every node
then, using this information just get the result by this crazy formula:
    ans[v] = ans[root] + n - 2*(childrencount[v]),
which i ensure u, have the underlying meaning but... you can reveal it by yourself! briefly, we just use the result for parent node to calculate the answer for the current node
*/
