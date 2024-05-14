package main

import "fmt"

func SchedulableDays(date1 []int, date2 []int) []int {
	var result []int //ngasih tau tipe variabel
	result = make([]int, 0) //bikin array
	for _, dateImam := range date1 {
		for _, datePermana := range date2 {
			if dateImam == datePermana {
				result = append(result, dateImam)
				break
			}
		}
	}
	return result // TODO: replace this
}

func main() {
	fmt.Println(SchedulableDays([]int{1, 2, 3, 4}, []int{3, 4, 5, 6}))
}