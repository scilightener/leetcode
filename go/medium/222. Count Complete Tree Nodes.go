package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left, right := root.Left, root.Right
	leftHeight, rightHeight := 1, 1
	for left != nil {
		left = left.Left
		leftHeight++
	}

	for right != nil {
		right = right.Right
		rightHeight++
	}

	if leftHeight == rightHeight {
		return pow(2, leftHeight) - 1
	}

	return 1 + countNodes(root.Left) + countNodes(root.Right)
}

func pow(a, b int) int {
	tmp := a
	for b > 1 {
		tmp *= a
		b--
	}

	return tmp
}
