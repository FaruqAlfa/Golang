package main

import "fmt"

func TicketPlayground(height, age int) int {
	var price int

	if age < 5 {
		price = -1
	} else {
		price = 10000

		if age >= 5 || height > 120 {
			price += 5000
		}

		if age >= 8 || height > 135 {
			price += 10000
		}

		if age >= 10 || height > 150 {
			price += 15000
		}

		if age >= 12 || height > 160 {
			price = 60000
		}

		if age > 12 {
			price = 100000
		}
	}

	return price // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(TicketPlayground(160, 11))
}
