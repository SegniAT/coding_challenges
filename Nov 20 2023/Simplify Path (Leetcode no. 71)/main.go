package main

import (
	"strings"
)

func main() {
	simplifyPath("/../hi/bruh/../hi")
}

type stack []string

func newStack() *stack {
	var s stack = []string{}
	return &s
}

func (s *stack) isEmpty() bool {
	return len(*s) == 0
}

func (s *stack) top() string {
	if s.isEmpty() {
		return "/"
	}

	length := len(*s)
	return (*s)[length-1]
}

func (s *stack) push(item string) {
	*s = append((*s), item)
}

func (s *stack) pop() string {
	if s.isEmpty() {
		return "/"
	}

	length := len(*s)
	last := (*s)[length-1]
	*s = (*s)[0 : length-1]
	return last
}

func simplifyPath(path string) string {
	paths := strings.Split(path, "/")
	s := newStack()

	for _, p := range paths {
		if p == ".." {
			s.pop()
		} else if p == "" || p == "." {
			continue
		} else {
			s.push(p)
		}
	}

	res := ""

	for !s.isEmpty() {
		res = "/" + s.pop() + res
	}

	if res == "" {
		res = "/"
	}

	return res
}
