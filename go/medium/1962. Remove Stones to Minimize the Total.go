package medium

import (
	"container/heap"
)

type any = interface{}
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func minStoneSum(piles []int, k int) int {
	s := 0
	for i, v := range piles {
		piles[i] = -v
		s += v
	}
	h := IntHeap(piles)
	heap.Init(&h)
	for ; k > 0; k-- {
		n := h[0] / 2
		s += n
		h[0] -= n
		heap.Fix(&h, 0)
	}
	return s
}

/*
complexity:
	time: O(N + KlogN) - N for heap.Init and KlogN for k fixes
	space: O(1) - everything in-place

the idea is to use maxHeap to know the maximum elemnt at each iteration through k
then just do the rules and update heap (extract floor and heap.Fix)
*/
