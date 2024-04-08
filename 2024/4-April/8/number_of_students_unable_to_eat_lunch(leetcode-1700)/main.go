package main

import "fmt"

func main() {
	students := []int{1, 1, 0, 0}
	sandwiches := []int{0, 1, 0, 1}

	fmt.Println(countStudents(students, sandwiches))
}

func countStudents(students []int, sandwiches []int) int {
	for len(students) > 0 {
		if students[0] == sandwiches[0] {
			students = students[1:]
			sandwiches = sandwiches[1:]
		} else {
			// change positions in queue
			switches := 0
			studentsLen := len(students)
			for students[0] != sandwiches[0] {
				if switches >= studentsLen {
					return studentsLen
				}
				first := students[0]
				students = students[1:]
				students = append(students, first)
				switches++
			}
		}
	}

	return len(students)
}
