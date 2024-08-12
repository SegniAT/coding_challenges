package main

import "container/heap"

func main() {
}

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	old := *h
	N := len(*h)
	last := old[N-1]
	*h = old[:N-1]
	return last
}

func (h *MinHeap) Peek() int {
	if h.Len() > 0 {
		return (*h)[0]
	}

	panic("heap is empty")
}

type KthLargest struct {
	k       int
	minHeap *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	res := &KthLargest{
		k:       k,
		minHeap: &MinHeap{},
	}

	heap.Init(res.minHeap)

	for _, num := range nums {
		res.Add(num)
	}

	return *res
}

func (this *KthLargest) Add(val int) int {
	if this.minHeap.Len() < this.k {
		heap.Push(this.minHeap, val)
	} else if this.minHeap.Peek() < val {
		heap.Pop(this.minHeap)
		heap.Push(this.minHeap, val)
	}

	return this.minHeap.Peek()
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
