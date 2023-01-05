package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxAncestorDiff(root *TreeNode) int {
	return maxDiff(root, root.Val, root.Val, -1)
}

func maxDiff(root *TreeNode, curMin, curMax, ans int) int {
	if root == nil {
		return ans
	}
	ans = max(abs(curMax-root.Val), abs(curMin-root.Val))
	curMin = min(curMin, root.Val)
	curMax = max(curMax, root.Val)
	left, right := -1, -1
	if root.Left != nil {
		left = maxDiff(root.Left, curMin, curMax, ans)
	}
	if root.Right != nil {
		right = maxDiff(root.Right, curMin, curMax, ans)
	}
	ans = max(ans, left)
	ans = max(ans, right)
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
