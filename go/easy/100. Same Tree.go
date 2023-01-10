package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == nil && q == nil
	}
	if p.Val != q.Val || !isSameTree(p.Left, q.Left) || !isSameTree(p.Right, q.Right) {
		return false
	}
	return true
}

/*
complexity:
	time: O(N) for dfs
	space: O(N) for dfs

what we want? we want the roots to be identical first,
and the left and right children to be identical too
that's what there's written: if at least one of these conditions is not true, return false
otherwise, true
don't forget to handle the nil-case!!
*/
