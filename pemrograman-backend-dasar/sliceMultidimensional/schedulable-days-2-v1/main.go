package main

import (
	"fmt"
	"sort"
)

func SchedulableDays(villager [][]int) []int {
	avaliableDate := make(map[int]int)

	//mengembalikan nilai kosong jika tidak ada warga desa
	if len(villager) == 0 {
		return []int{}
	}

	//menghitung frekuensi setiap tanggal
	for _, schedule := range villager {
		for _, date := range schedule {
			avaliableDate[date]++
		}
	}

	//mencari tanggal yang tersedia
	var schedulableDays []int
	for date, frequency := range avaliableDate {
		if frequency == len(villager) {
			schedulableDays = append(schedulableDays, date)
		}
	}

	//mengurutkan tanggal
	sort.Ints(schedulableDays)

	return schedulableDays // TODO: replace this
}

func main() {
	var villager = [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
		{3, 5, 7, 9, 11, 13, 15, 17, 19, 21},
		{5, 7, 9, 11, 13, 15, 17, 19, 21, 23},
	}
	fmt.Println(SchedulableDays(villager))
	// TODO: answer here
}
