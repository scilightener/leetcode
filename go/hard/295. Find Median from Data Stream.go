package hard

import "container/heap"

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type MedianFinder struct {
	minHeap *IntHeap
	maxHeap *IntHeap
}

func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: new(IntHeap),
		maxHeap: new(IntHeap),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if (this.minHeap.Len()+this.maxHeap.Len())%2 == 0 {
		heap.Push(this.maxHeap, -num)
		heap.Push(this.minHeap, -heap.Pop(this.maxHeap).(int))
	} else {
		heap.Push(this.minHeap, num)
		heap.Push(this.maxHeap, -heap.Pop(this.minHeap).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if (this.minHeap.Len()+this.maxHeap.Len())%2 == 0 {
		return (float64((*this.minHeap)[0]) - float64((*this.maxHeap)[0])) / 2
	}
	return float64((*this.minHeap)[0])
}
