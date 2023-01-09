package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) []int {
	ans := []int{}
	dfs(root, &ans)
	return ans
}

func dfs(node *TreeNode, path *[]int) {
	if node == nil {
		return
	}
	*path = append(*path, node.Val)
	dfs(node.Left, path)
	dfs(node.Right, path)
}

/*
complexity:
	time: O(N) just dfs
	space: O(N) dfs worst case recursion stack

in dfs: first, visit the root, then visit left subtree and right subtree, everything is clear
if node is nil, then skip
*/
