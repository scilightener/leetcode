package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leaves1, leaves2 := make([]int, 0), make([]int, 0)
	dfs(root1, &leaves1)
	dfs(root2, &leaves2)
	if len(leaves1) != len(leaves2) {
		return false
	}
	for i, val := range leaves1 {
		if leaves2[i] != val {
			return false
		}
	}

	return true
}

func dfs(root *TreeNode, leaves *[]int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		*leaves = append(*leaves, root.Val)
	}
	dfs(root.Left, leaves)
	dfs(root.Right, leaves)
}
