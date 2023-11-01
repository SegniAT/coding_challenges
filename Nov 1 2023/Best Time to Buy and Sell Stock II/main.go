package main

import (
	"fmt"
	"math"
)

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	mq := createMonotonicQueue(len(input))
	profit := 0

	for i := len(input) - 1; i > -1; i-- {
		item := Item{
			index: i,
			value: input[i],
		}
		profit += mq.push(item)
	}
}

func maxProfit(prices []int) int {
	profit := 0.0

	for i := 0; i < len(prices)-1; i++ {
		profit += math.Max(0, float64(prices[i+1])-float64(prices[i]))
	}

	return int(profit)
}

type Item struct {
	value int
	index int
}

type MonotonicQueue struct {
	queue []Item
}

func createMonotonicQueue(size int) *MonotonicQueue {
	mq := MonotonicQueue{}
	mq.queue = make([]Item, size)
	return &mq
}

func (mq *MonotonicQueue) isEmpty() bool {
	return len(mq.queue) == 0
}

func (mq *MonotonicQueue) getLast() Item {
	length := len(mq.queue)
	if length < 0 {
		panic("getLast() on empty array")
	}
	return mq.queue[length-1]
}

func (mq *MonotonicQueue) removeLast() {
	length := len(mq.queue)
	if length < 0 {
		panic("removeFirst() on empty array")
	}
	mq.queue = mq.queue[:length-1]
}

func (mq *MonotonicQueue) push(item Item) int {
	profit := 0

	for !mq.isEmpty() && mq.getLast().value < item.value {
		mq.removeLast()
	}

	if !mq.isEmpty() && mq.getLast().index-item.index == 1 {
		profit += mq.getLast().value - item.value
	}

	mq.queue = append(mq.queue, item)
	fmt.Println("QUEUE: ", mq.queue)

	return profit
}

func maxProfitMQ(prices []int) int {
	mq := createMonotonicQueue(len(prices))
	profit := 0

	for i := len(prices) - 1; i > -1; i-- {
		item := Item{
			index: i,
			value: prices[i],
		}
		profit += mq.push(item)
	}
	return profit
}
