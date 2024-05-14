package main

import "fmt"

func BMICalculator(gender string, height int) float64 {

	var bmi float64

	if gender == "laki-laki" {
		bmi = (float64(height - 100) - ((float64(height - 100) * 0.1)))
	}else{
		bmi = (float64(height - 100) - ((float64(height - 100) * 0.15)))
	}

	return bmi // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(BMICalculator("laki-laki", 170))
	fmt.Println(BMICalculator("perempuan", 165))
}
