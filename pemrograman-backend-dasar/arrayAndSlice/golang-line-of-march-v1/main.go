package main

import (
	"fmt"
	"sort"
)

func Sortheight(height []int) []int {
	sort.Ints(height)	
	return height // TODO: replace this
}


func main() {
	height := []int{180, 164, 170, 171, 172, 173}
	fmt.Println(Sortheight(height))
}