package main

import (
	"fmt"
	"math"
)

func main() {
}

func maxProfitNaive(prices []int) int {
	maxProf := math.Inf(-1)
	for i := 0; i < len(prices); i++ {
		for j := i; j < len(prices); j++ {
			maxProf = math.Max(maxProf, float64(prices[j])-float64(prices[i]))
		}
	}
	return int(maxProf)
}

type Item struct {
	index int
	value int
}

type Queue[T interface{}] struct {
	items []T
}

func newQueue[T interface{}](size int) *Queue[T] {
	q := Queue[T]{}
	q.items = make([]T, size)
	return &q
}

func (q *Queue[T]) isEmpty() bool {
	return len(q.items) == 0
}

func (q *Queue[T]) push(item T) {
	q.items = append(q.items, item)
}

func (q *Queue[T]) removeLast() T {
	lastIndex := len(q.items) - 1
	item := q.items[lastIndex]
	q.items = q.items[:lastIndex]
	return item
}

func (q *Queue[T]) peekLast() T {
	lastIndex := len(q.items) - 1
	return q.items[lastIndex]
}

func (q *Queue[T]) peekFirst() T {
	return q.items[0]
}

func maxProfitMonotonicQueue(prices []int) int {
	queue := newQueue[Item](len(prices))
	profits := make([]int, len(prices))

	for i := len(prices) - 1; i > -1; i-- {
		pushDecreasingMonotonicQueue(queue, Item{index: i, value: prices[i]}, &profits)
	}

	fmt.Println(profits)
	var maxProf float64 = 0
	for _, profit := range profits {
		maxProf = math.Max(maxProf, float64(profit))
	}
	return int(maxProf)
}

func pushDecreasingMonotonicQueue(q *Queue[Item], item Item, profits *[]int) {
	for !q.isEmpty() && item.value >= q.peekLast().value {
		q.removeLast()
	}

	if !q.isEmpty() {
		(*profits)[item.index] = q.peekFirst().value - item.value
	} else {
		(*profits)[item.index] = 0
	}

	q.push(item)
}

func maxProfitSlidingWindow(prices []int) int {
	left, right, maxProfit := 0, 1, 0.0

	for right < len(prices) {
		if prices[left] < prices[right] {
			profit := prices[right] - prices[left]
			maxProfit = math.Max(maxProfit, float64(profit))
		} else {
			left = right
		}
		right++
	}

	return int(maxProfit)
}
