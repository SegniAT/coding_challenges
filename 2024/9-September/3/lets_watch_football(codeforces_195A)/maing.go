package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := strings.Split(scanner.Text(), " ")
	a, _ := strconv.Atoi(line[0])
	b, _ := strconv.Atoi(line[1])
	c, _ := strconv.Atoi(line[2])

	fmt.Println(minTime(a, b, c))
}

/*
NOTE:
	a - data needed to watch 1 second
 	b - our internet speed
	c - video length in seconds
*/

func minTime(a, b, c int) int {
	// NOTE: using just a*c instead of a*c/b to not underestimate the upper bound...integer division might result in lower than the acual answer
	minTime, maxTime := 0, a*c

	for minTime < maxTime {
		mid := minTime + (maxTime-minTime)/2

		// NOTE: data downloaded after time waited (mid) till the video ends (c)
		dataDownloaded := b * (mid + c)
		totalDataNeeded := a * c // Total data needed to watch the entire video

		if dataDownloaded < totalDataNeeded {
			minTime = mid + 1
		} else {
			maxTime = mid
		}
	}

	return minTime
}
