package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	nextNode := node.Next
	for ; nextNode.Next != nil; nextNode = node.Next {
		node.Val = nextNode.Val
		node = nextNode
	}
	node.Val = nextNode.Val
	node.Next = nil
}
