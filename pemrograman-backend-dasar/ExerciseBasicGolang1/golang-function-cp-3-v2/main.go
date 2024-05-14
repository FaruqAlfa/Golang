package main

import (
	"fmt"
)

func WhichOneIsShorter(name1 string, name2 string) string {
	if len(name1) < len(name2) {
		return name1
	} else if len(name1) > len(name2) {
		return name2
	} else {
		if name1 < name2 {
			return name1
		} else {
			return name2
		}
	}
}

func FindShortestName(names string) string {
	curName := ""
	minName := "asudyrhuqwagfhjawfjhguyweagrfyuwdgsfjkdhgfuywkeghfjad"
	for _, c := range names {
		if string(c) == ";" || string(c) == "," || string(c) == " " {
			minName = WhichOneIsShorter(minName, curName)
			curName = ""
		} else {
			curName += string(c)
		}
	}
	if curName != "" {
		minName = WhichOneIsShorter(minName, curName)
	}
	return minName // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(FindShortestName("Hanif Joko Tio Andi Budi Caca Hamdan")) // "Tio"
	fmt.Println(FindShortestName("Budi;Tia;Tio"))                         // "Tia"
}
