package medium

import "container/heap"

type WordHeapItem struct {
	word string
	cnt  int
}

type WordHeap []WordHeapItem

func (h WordHeap) Len() int {
	return len(h)
}

func (h WordHeap) Less(i, j int) bool {
	if h[i].cnt != h[j].cnt {
		return h[i].cnt > h[j].cnt
	}
	return h[i].word < h[j].word
}

func (h WordHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *WordHeap) Push(x interface{}) {

	*h = append(*h, x.(WordHeapItem))
}

func (h *WordHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

var counter map[string]int

func topKFrequent(words []string, k int) []string {
	counter = make(map[string]int)
	h := &WordHeap{}
	for _, word := range words {
		if _, ok := counter[word]; !ok {
			counter[word] = 0
		}
		counter[word]++
	}

	for key, val := range counter {
		*h = append(*h, WordHeapItem{key, val})
	}

	heap.Init(h)
	res := make([]string, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(h).(WordHeapItem).word
	}
	return res
}
