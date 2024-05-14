package main

import (
	"fmt"
)

func MapToSlice(mapData map[string]string) [][]string {
	var result [][]string

    for key, value := range mapData {
        result = append(result, []string{key, value})
    }
	return result // TODO: replace this
}

func main() {
	fmt.Println(MapToSlice(map[string]string{
		"hello": "world",
		"John":  "Doe",
		"Age":   "14",
		"foo":   "33",
		"bar":   "44",
		"baz":   "55",
	}))
}
