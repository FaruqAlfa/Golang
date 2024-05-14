package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	Hour   int
	Minute int
}


func ChangeToStandartTime(time any) string {
	var hour int
	var minute int

	switch val := time.(type) {
	case string:
		// val := time.(string)
		tokens := strings.Split(val, ":")
		if len(tokens) != 2 {
			return "Invalid input"
		}
		if tokens[0] == "" {
			return "Invalid input"
		}
		hour, _ = strconv.Atoi(tokens[0])
		if tokens[1] == ""{
			return "Invalid input"
		}
		minute, _ = strconv.Atoi(tokens[1])

	case []int:
		// val := time.([]int)
		if len (val) != 2 {
			return "Invalid input"
		}

		hour = val[0]
		minute = val[1]
	case map[string]int:
		// val := time.(map[string]int)
		var oke bool
		hour, oke = val["hour"]
		if !oke {
			return "Invalid input"
		}
		minute, oke = val["minute"]
		if !oke {
			return "Invalid input"
		}

	case Time:
		// val := time.(Time)
		hour = val.Hour
		minute = val.Minute
	}
	fmt.Println("hour: ", hour)
	fmt.Println("minute: ", minute)
	amPm := "AM"
	if hour >= 12 {
		amPm = "PM"
	}
	if hour > 12 {
		hour -= 12
	}

	return fmt.Sprintf("%02d:%02d %s", hour, minute, amPm) // TODO: replace this
}

func main() {
	fmt.Println(ChangeToStandartTime("16:00"))
	fmt.Println(ChangeToStandartTime([]int{16, 0}))
	fmt.Println(ChangeToStandartTime(map[string]int{"hour": 16, "minute": 0}))
	fmt.Println(ChangeToStandartTime(Time{16, 0}))
}
