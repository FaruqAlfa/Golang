package main

import (
	"fmt"
	"strings"
	"strconv"
)

// TODO: answer here

func ChangeOutput(data []string) map[string][]string {
	result := make (map[string][]string)
	for _, d := range data {
		fmt.Println("d", d)
		tokens := strings.Split(d, "-")
		fmt.Printf("tokens %#v\n", tokens)

		label := tokens[0]
		index, _ := strconv.Atoi(tokens[1])
		fristOrLast := tokens[2]
		value := tokens[3]

		if _, ok := result[label]; !ok {
			result[label] = make([]string, 0)
		}

		if fristOrLast == "frist" {
			if index >= len (result[label]) {
				result[label] = append(result[label], value)
			}else {
				result[label][index] = value + result[label][index]
			}
		}else {
			if index >= len (result[label]) {
				result[label] = append(result[label], value)
			}else {
				result[label][index] = result[label][index] + " " + value
			}
		}
		fmt.Println()
	}
	fmt.Printf("result %#v\n", result)
	return result // TODO: replace this
}

// bisa digunakan untuk melakukan debug
func main() {
	data := []string{"account-0-first-John", "account-0-last-Doe", "account-1-first-Jane", "account-1-last-Doe", "address-0-first-Jaksel", "address-0-last-Jakarta", "address-1-first-Bandung", "address-1-last-Jabar", "phone-0-first-081234567890", "phone-1-first-081234567891"}
	res := ChangeOutput(data)

	fmt.Println(res)
}
