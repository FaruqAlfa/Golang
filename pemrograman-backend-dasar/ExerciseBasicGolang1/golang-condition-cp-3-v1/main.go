package main

import "fmt"

func GetPredicate(math, science, english, indonesia int) string {

	var peringkat string
	var score int = (math + science + english + indonesia) / 4

	switch {
	case score == 100:
		peringkat = "Sempurna"
	case score >= 90 && score < 100:
		peringkat = "Sangat Baik"
	case score >= 80 && score < 90:
		peringkat = "Baik"
	case score >= 70 && score < 80:
		peringkat = "Cukup"
	case score >= 60 && score < 70:
		peringkat = "Kurang"
	default:
		peringkat = "Sangat kurang"
	}

	return peringkat // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(GetPredicate(50, 80, 100, 60))
}
