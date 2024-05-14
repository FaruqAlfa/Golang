package main

import "fmt"

func SlurredTalk(words *string) {
	result := ""
	for _, v := range *words {
		// fmt.Println(string(v))
		if v == 'S' || v == 'R' || v == 'Z' {
			result += "L"
		}else if  v == 's' || v == 'r' || v == 'z'{
			result += "l"
		}else {
			result += string(v)
		}
	}
	// fmt.Println(result)
	*words = result
	// TODO: answer here
}

func main() {
	// bisa dicoba untuk pengujian test case
	var words string = "Semangat pagi semuanya, semoga sehat selalu. Sehingga selalu bisa senyum"
	SlurredTalk(&words)
	fmt.Println(words)
}
