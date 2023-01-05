package medium

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxProduct(root *TreeNode) int {
	var list []int
	sum := getSum(root, &list)

	res := 0
	for _, v := range list {
		if sum/2-v < 0 {
			v = sum - v
		}
		if sum/2-v < sum/2-res {
			res = v
		}
	}
	return res * (sum - res) % 1000000007
}

func getSum(root *TreeNode, list *[]int) int {
	if root == nil {
		return 0
	}
	res := root.Val + getSum(root.Left, list) + getSum(root.Right, list)
	*list = append(*list, res)
	return res
}
