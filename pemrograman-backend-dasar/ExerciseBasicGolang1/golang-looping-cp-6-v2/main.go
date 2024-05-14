package main

import (
	"fmt"
	"strconv"
)

func BiggestPairNumber(numbers int) int {
	var pair = 0
	var biggest = 0

	numStr := strconv.Itoa(numbers)

	for i := 1; i < len(numStr); i++ {
		num1, _ := strconv.Atoi(string(numStr[i]))
		num2, _ := strconv.Atoi(string(numStr[i-1]))

		if num1+num2 > biggest {
			biggest = num1 + num2

			pairNum, _ := strconv.Atoi(string(numStr[i-1]) + string(numStr[i]))
			pair = pairNum
		}

	}

	return pair // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BiggestPairNumber(11223344))
}
