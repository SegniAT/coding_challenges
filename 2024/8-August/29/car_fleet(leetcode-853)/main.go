package main

import (
	"fmt"
	"sort"
)

func main() {
	target := 12
	position := []int{10, 8, 0, 5, 3}
	speed := []int{2, 4, 1, 1, 3}

	target = 100
	position = []int{0, 2, 4}
	speed = []int{4, 2, 1}

	fmt.Println(carFleet(target, position, speed))
}

type Car struct {
	pos   int
	speed int
}

func (car Car) timeToReach(target int) float64 {
	if target < car.pos {
		panic("target position too low")
	}

	return (float64(target) - float64(car.pos)) / float64(car.speed)
}

// just make sure that there are no faster cars behind you...faster meaning less time to reach target
// so keep a decreasing monotonic stack...no car is faster than cars in front

func carFleet(target int, position []int, speed []int) int {
	N := len(position)
	cars := make([]Car, N)

	for i := range N {
		cars[i] = Car{
			pos:   position[i],
			speed: speed[i],
		}
	}

	sort.Slice(cars, func(i, j int) bool {
		return cars[i].pos < cars[j].pos
	})

	fmt.Println("sorted: ", cars)

	stack := []Car{}

	for _, car := range cars {
		for len(stack) > 0 && stack[len(stack)-1].timeToReach(target) <= car.timeToReach(target) {
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, car)
	}

	return len(stack)
}
