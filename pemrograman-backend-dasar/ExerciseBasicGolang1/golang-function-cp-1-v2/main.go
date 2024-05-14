package main

import (
	"fmt"
)

func DateFormat(day, month, year int) string {

	var formatted_month string

	if month == 1 {
		formatted_month = "January"
	} else if month == 2 {
		formatted_month = "February"
	} else if month == 3 {
		formatted_month = "March"
	} else if month == 4 {
		formatted_month = "April"
	} else if month == 5 {
		formatted_month = "May"
	} else if month == 6 {
		formatted_month = "June"
	} else if month == 7 {
		formatted_month = "July"
	} else if month == 8 {
		formatted_month = "August"
	} else if month == 9 {
		formatted_month = "September"
	} else if month == 10 {
		formatted_month = "October"
	} else if month == 11 {
		formatted_month = "November"
	}else if month == 12 {
		formatted_month = "December"
	}

	var formatted_day string

	if day < 10 {
		formatted_day = "0" + fmt.Sprintf("%d", day)
	} else {
		formatted_day = fmt.Sprintf("%d", day)
	}

	return formatted_day + "-" + formatted_month + "-" + fmt.Sprintf("%d", year) // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(DateFormat(1, 1, 2012))

}
