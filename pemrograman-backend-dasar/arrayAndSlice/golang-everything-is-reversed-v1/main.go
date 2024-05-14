package main

import "fmt"

func ReverseData(arr [5]int) [5]int {
	reversed := [5]int{}
	for i := 0; i < len(arr); i++ {
		reversed[len(arr)-1-i] = reversedDigits(arr[i])
	}
	
	return reversed // TODO: replace this
}

func reversedDigits(number int) int {
	reversed := 0
	for number != 0 {
		reversed = reversed * 10 + number % 10
		number = number / 10
	}

	return reversed
}

func main() {
	digit := [5]int{21, 32, 43, 54, 65}
	fmt.Println(ReverseData(digit))
}