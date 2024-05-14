package main

import (
	"fmt"
	"strings"
)

func FindSimilarData(input string, data ...string) string {
	result := ""
	yangPertama := true
	for _, v := range data {
		// fmt.Println(v)
		apakahMengnandungInput := strings.Contains(v, input)
		// fmt.Println(apakahMengnandungInput)

		if apakahMengnandungInput {
			if yangPertama {
				result += v
				yangPertama = false
			}else{
				result += "," +  v
			}
			
		}
	}
	// fmt.Println(result)
	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindSimilarData("iphone", "laptop", "iphone 13", "iphone 12", "iphone 12 pro"))
}
