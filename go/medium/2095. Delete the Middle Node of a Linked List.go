package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil {
		return nil
	}
	slow, fast := head, head
	var prev *ListNode = nil
	for ; fast != nil && fast.Next != nil; fast = fast.Next.Next {
		prev, slow = slow, slow.Next
	}
	prev.Next = slow.Next
	return head
}
