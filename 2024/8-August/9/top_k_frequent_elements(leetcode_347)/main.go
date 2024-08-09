package main

import "container/heap"

func main() {
}

func topKFrequent(nums []int, k int) []int {
	pq := &PriorityQueue{}
	heap.Init(pq)

	countCache := make(map[int]int)
	for _, num := range nums {
		countCache[num]++
	}

	for key, val := range countCache {
		heap.Push(pq, &Item{value: key, priority: val})
	}

	res := []int{}

	for range k {
		res = append(res, heap.Pop(pq).(*Item).value)
	}

	return res
}

type Item struct {
	value    int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority > pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[:n-1]
	return item
}
