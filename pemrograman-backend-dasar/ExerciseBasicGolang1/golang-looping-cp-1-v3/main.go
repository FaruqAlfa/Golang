package main

import "fmt"

func CountingNumber(n int) float64 {
	var total float64

	for i := 1.0; i <= float64(n); i += 0.5 {
		total += float64(i)
	}

	return total // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingNumber(10))
}
