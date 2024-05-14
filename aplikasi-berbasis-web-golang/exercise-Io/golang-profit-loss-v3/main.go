package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Readfile(path string) ([]string, error) {

	content, err := os.ReadFile(path)
	if err != nil{
		panic(err)
	}
	txt := string(content)

	lines := strings.Split(txt, "\n")
	if len(lines) == 0 {
		return []string{}, nil
	}else if len(lines) == 1 && lines[0] == "" {
		return []string{}, nil
		
	}
	return lines, nil // TODO: replace this
}

func CalculateProfitLoss(data []string) string {
	var profit int
	lastDate := ""

	for _, line := range data {
		datas := strings.Split(line, ";")
		date := datas[0]
		lastDate = date
		transaction := datas[1]
		amount, _ := strconv.Atoi(datas[2])

		if transaction == "income" {
			profit += amount
		} else if transaction == "expense" {
			profit -= amount
		}
	}

	if profit >= 0 {
		return fmt.Sprintf("%s;profit;%d", lastDate, profit)
	} else {
		return fmt.Sprintf("%s;loss;%d", lastDate, profit*-1)
	}
	// TODO: replace this
}

func main() {
	// bisa digunakan untuk pengujian
	datas, err := Readfile("transactions.txt")
	if err != nil {
		panic(err)
	}

	result := CalculateProfitLoss(datas)
	fmt.Println(result)
}
