package hard

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var ans int

func maxPathSum(root *TreeNode) int {
	ans = -1001
	dfs(root)
	return ans
}

func dfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := max(dfs(root.Left), 0), max(dfs(root.Right), 0)
	ans = max(ans, left+root.Val+right)
	return root.Val + max(left, right)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
