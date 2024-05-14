package main

import "fmt"

func CountingLetter(text string) int {
	// unreadable letters = R, S, T, Z
	var count int
	for i := 0; i < len(text); i++ {
		if text[i] == 'R' || text[i] == 'r' || text[i] == 'S' || text[i] == 's' || text[i] == 'T' || text[i] == 't' || text[i] == 'Z' || text[i] == 'z' {
			count++
		}
	}
	return count // TODO: replace this

}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(CountingLetter("Semangat"))
}
