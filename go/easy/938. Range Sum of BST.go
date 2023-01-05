package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}
	return dfs(root, low, high, 0)
}

func dfs(root *TreeNode, low, high, count int) int {
	if low <= root.Val && root.Val <= high {
		count += root.Val
	}
	if root.Left != nil {
		count = dfs(root.Left, low, high, count)
	}
	if root.Right != nil {
		count = dfs(root.Right, low, high, count)
	}
	return count
}
