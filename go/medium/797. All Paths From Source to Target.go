package _go

func allPathsSourceTarget(graph [][]int) [][]int {
	paths := [][]int{}
	stack := [][]int{{0}}
	for len(stack) > 0 {
		path := stack[len(stack)-1]
		u := path[len(path)-1]
		stack = stack[:len(stack)-1]

		if u == len(graph)-1 {
			paths = append(paths, path)
		}
		for _, v := range graph[u] {
			stack = append(stack, append(path[:len(path):len(path)], v))
		}
	}
	return paths
}

/*
complexity:
	time: O(E) dfs on edges of the graph
	space: O(???) we store all the possible paths so dunno...

the idea is to save all the possible paths, provided by something like dfs
then just go through them and find the desired ones
*/
