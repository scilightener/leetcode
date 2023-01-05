package medium

type ListNode struct {
	Val  int
	Next *ListNode
}

func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return head
	}
	evenHead, lastOdd, cur, isOdd := head.Next, head, head, true
	for cur.Next != nil {
		next := cur.Next
		cur.Next = cur.Next.Next
		cur, isOdd = next, !isOdd
		if isOdd {
			lastOdd = cur
		}
	}
	lastOdd.Next = evenHead
	return head
}
