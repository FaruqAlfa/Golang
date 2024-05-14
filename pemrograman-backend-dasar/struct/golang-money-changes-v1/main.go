package main

import "fmt"

type Product struct {
	Name  string
	Price int
	Tax   int
}

func ExchangeCoin(amount int) []int {
	coin := make ([]int, 0)

	for amount > 0 {
		if amount >= 1000 {
			coin = append(coin, 1000)
			amount = amount - 1000
		} else if amount >= 500 {
			coin = append(coin, 500)
			amount = amount - 500
		} else if amount >= 200 {
			coin = append(coin, 200)
			amount = amount - 200
		} else if amount >= 100 {
			coin = append(coin, 100)
			amount = amount - 100
		} else if amount >= 50 {
			coin = append(coin, 50)
			amount = amount - 50
		} else if amount >= 20 {
			coin = append(coin, 20)
			amount = amount - 20
		} else if amount >= 10 {
			coin = append(coin, 10)
			amount = amount - 10
		} else if amount >= 5 {
			coin = append(coin, 5)
			amount = amount - 5
		} else if amount >= 1 {
			coin = append(coin, 1)
			amount = amount - 1
		}
	}

	return coin// TODO: replace this
}

func MoneyChanges(amount int, products []Product) []int {
	cost := 0

	for _, product := range products{
		cost += product.Price + product.Tax
	}

	change := amount - cost
	return ExchangeCoin(change) // TODO: replace this
}

func main(){
	products := []Product{
		{Name: "kacang", Price: 3000, Tax: 0},
		{Name: "jahe", Price: 5000, Tax: 500},
		{Name: "kencur", Price: 10000, Tax: 0},
	}

	changes := MoneyChanges(20000, products)
	fmt.Println(changes)
}
