package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	str, err := loadInput()
	if err != nil {
		log.Fatal("error loading input: ", err)
	}

	str += "x"

	prevChar, count := str[0], 0

	for i := 0; i < len(str); i++ {
		if currChar := str[i]; currChar == prevChar {
			count++
		} else {
			if count >= 7 {
				fmt.Println("YES")
				return
			}
			prevChar = currChar
			count = 1
		}
	}

	fmt.Println("NO")
}

func loadInput() (string, error) {
	scanner := bufio.NewScanner(os.Stdin)

	// test cases
	scanner.Scan()
	return scanner.Text(), nil
}
