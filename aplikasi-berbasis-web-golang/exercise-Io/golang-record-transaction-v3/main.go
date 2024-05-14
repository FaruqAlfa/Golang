package main

import (
	// "context"
	// "errors"
	"fmt"
	"os"
	"sort"
	"strings"
)

type Transaction struct {
	Date   string
	Type   string
	Amount int
}

func RecordTransactions(path string, transactions []Transaction) error {
	if len(transactions) == 0 {
		return nil
	}

	sort.Slice(transactions, func(i, j int) bool {
		return transactions[i].Date < transactions[j].Date
	})
	lastDate := transactions[0].Date
	money := 0
	output := make([]string, 0)

	for _, transaction := range transactions {
		if transaction.Date == lastDate {
			if transaction.Type == "income" {
				money += transaction.Amount
			} else {
				money -= transaction.Amount
			}
		} else {
			if money < 0 {
				output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, money*-1))
			}else {
				output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
			}

			money = 0
			if transaction.Type == "income" {
				money += transaction.Amount
			} else {
				money -= transaction.Amount
			}

			lastDate = transaction.Date
		}
	}

	if money < 0 {
		output = append(output, fmt.Sprintf("%s;expense;%d", lastDate, money*-1))
	}else {
		output = append(output, fmt.Sprintf("%s;income;%d", lastDate, money))
	}

	
	content, err := os.Create(path)
	if err != nil {
		return err
	}
	defer content.Close()

	_, err = content.WriteString(strings.Join(output, "\n"))
	if err != nil {
		return err
	}

	return nil // TODO: replace this
}


func main() {
	// bisa digunakan untuk pengujian test case
	var transactions = []Transaction{
		{"01/01/2021", "expense", 100000},
					{"01/01/2021", "expense", 1000},
					{"02/01/2021", "expense", 3424},
					{"02/01/2021", "expense", 42000},
					{"03/01/2021", "expense", 22321},
					{"04/01/2021", "expense", 223200},
					{"02/01/2021", "expense", 2300},
					{"05/01/2021", "expense", 2213},
					{"06/01/2021", "expense", 4545},
					{"07/01/2021", "expense", 55500},
					{"08/01/2021", "expense", 200000},
					{"10/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 10000},
					{"12/01/2021", "expense", 55500},
					{"13/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 55500},
					{"02/01/2021", "expense", 10000},
					{"14/01/2021", "expense", 20000},
					{"11/01/2021", "expense", 20000},
					{"15/01/2021", "expense", 10000},
					{"16/01/2021", "expense", 20000},
					{"02/01/2021", "expense", 55500},
					{"17/01/2021", "expense", 10000},
					{"06/01/2021", "expense", 20000},
					{"18/01/2021", "expense", 10000},
					{"03/01/2021", "expense", 20000},
					{"04/01/2021", "expense", 10000},
					{"19/01/2021", "expense", 55500},
					{"20/01/2021", "expense", 55500},
					{"21/01/2021", "expense", 10000},
					{"22/01/2021", "expense", 10000},
					{"23/01/2021", "expense", 10000},
					{"24/01/2021", "expense", 10000},
	}

	err := RecordTransactions("transactions.txt", transactions)
	if err != nil {
		panic(err)
	}

	fmt.Println("Success")
}
