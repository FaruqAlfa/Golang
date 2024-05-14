package main

import (
	"fmt"
)
func CountProfit(data [][][2]int) []int {
	resultMap := make(map[int]int)
	for _, cabang := range data {
		fmt.Println("Cabang", cabang)
		for bulanKe, bulan := range cabang {
			fmt.Println("Bulan", bulan)

			income := bulan[0]
			expense := bulan[1]

			fmt.Println("Income", income)
			fmt.Println("Expense", expense)

			profit := income - expense
			resultMap[bulanKe + 1] += profit
		}
		fmt.Println()
	}
	numOfBulan := 0
	for k := range resultMap {
		if k > numOfBulan {
			numOfBulan = k
		}
	}
	result := make([]int, numOfBulan)

	for k, v := range resultMap {
		result[k-1] = v
	}
	return result // TODO: replace this
}

func main (){
	fmt.Println(CountProfit([][][2]int{
		{{1000, 500}, {500, 200}},
		{{1200, 200}, {1000, 800}},
		{{500, 100}, {700, 100}},
}))
}


