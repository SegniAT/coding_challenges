package main

import (
	"fmt"
	"math"
)

func main() {
}

type Queue struct {
	q []string
}

func (queue *Queue) enque(val string) {
	queue.q = append(queue.q, val)
}

func (queue *Queue) deque() string {
	if len(queue.q) > 0 {
		first := queue.q[0]
		queue.q = queue.q[1:]
		return first
	}
	return ""
}

func (queue *Queue) size() int {
	return len(queue.q)
}

func NewQueue() *Queue {
	return &Queue{
		q: []string{},
	}
}

func openLock(deadends []string, target string) int {
	deadendsSet := make(map[string]bool)
	visited := make(map[string]bool)
	for _, deadend := range deadends {
		deadendsSet[deadend] = true
	}

	queue := NewQueue()
	queue.enque("0000")

	steps := 0

	for queue.size() > 0 {
		size := queue.size()

		for range size {
			curr := queue.deque()
			if visited[curr] == true {
				continue
			} else {
				visited[curr] = true
			}

			if curr == target {
				return steps
			}

			if _, ok := deadendsSet[curr]; ok {
				continue
			}

			nextPatterns := generateLockPattern(curr)
			for _, pattern := range nextPatterns {
				queue.enque(pattern)
			}
		}

		steps++
	}

	return -1
}

func generateLockPattern(curr string) []string {
	res := []string{}

	for i := range curr {
		// increase
		char := curr[i]
		if char == '9' {
			char = '0'
		} else {
			char++
		}

		newStr := fmt.Sprintf("%s%s%s", curr[0:i], string(char), curr[i+1:])
		res = append(res, newStr)

		// decrease
		char = curr[i]
		if char == '0' {
			char = '9'
		} else {
			char--
		}

		newStr = fmt.Sprintf("%s%s%s", curr[0:i], string(char), curr[i+1:])
		res = append(res, newStr)
	}

	return res
}
