package medium

import (
	"container/heap"
	"sort"
)

type any = interface{}
type Item struct {
	begin    int
	duration int
	index    int
}
type Heap []Item

func (h Heap) Len() int      { return len(h) }
func (h Heap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h Heap) Less(i, j int) bool {
	if h[i].duration == h[j].duration {
		return h[i].index < h[j].index
	}
	return h[i].duration < h[j].duration
}

func (h *Heap) Push(x any) {
	item := x.(Item)
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}

func getOrder(tasks [][]int) []int {
	ans := make([]int, 0)
	sorted := make([]Item, len(tasks))
	for i, task := range tasks {
		sorted[i] = Item{task[0], task[1], i}
	}
	sort.Slice(sorted, func(i, j int) bool {
		return sorted[i].begin < sorted[j].begin
	})
	h := &Heap{}
	procTime, i := sorted[0].begin, 0
	for len(*h) > 0 || i < len(sorted) {
		for i < len(sorted) && sorted[i].begin <= procTime {
			heap.Push(h, sorted[i])
			i++
		}
		if len(*h) == 0 && i < len(sorted) {
			procTime = sorted[i].begin
			continue
		}
		t := heap.Pop(h).(Item)
		procTime += t.duration
		ans = append(ans, t.index)
	}
	return ans
}

/*
complexity:
	time: O(NlogN) for sorting + N times we push items to heap and pop them
	space: O(N) for heap

idea:
	first, ensure that all the tasks are sorted by their beginning time
	then, create heap for the available tasks at the current moment (procTime)
	push all the tasks that are already available for processor to perform (begin <= procTime)
	pop tasks from heap and execute them
	if none tasks, encrease procTime so that at least one task becomes available
*/
