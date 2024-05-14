package main

import (
	"fmt"
	"strconv"
	"strings"
)

func PopulationData(data []string) []map[string]interface{} {
	result := []map[string]interface{}{}
	for _, data := range data {
		currentResult := make(map[string]interface{})
		fmt.Println(data)
		tokens := strings.Split(data, ";")
		name := tokens[0]
		age, _ := strconv.Atoi(tokens[1])
		address := tokens[2]
		height, _ := strconv.ParseFloat(tokens[3], 64)
		isMarried, _:= strconv.ParseBool (tokens[4])
		fmt.Println("name: ", name)
		fmt.Println("age: ", age)
		fmt.Println("address: ", address)
		fmt.Println("height: ", height)
		fmt.Println("isMarried: ", isMarried)
		fmt.Println()

		currentResult["name"] = name
		currentResult["age"] = age
		currentResult["address"] = address
		if tokens[3] != "" {
			currentResult["height"] = height
		}
		if tokens[4] != "" {
			currentResult["isMarried"] = isMarried
		}
		result = append(result, currentResult)
	}
	return result // TODO: replace this
}

func main() {
	data := []string{
		"Budi;23;Jakarta;160.42;true",
		"Joko;30;Bandung;160;true",
		"Susi;25;Bogor;165.42;false",
	}

	res := PopulationData(data)
	fmt.Println(res)
} // TODO: answer here
