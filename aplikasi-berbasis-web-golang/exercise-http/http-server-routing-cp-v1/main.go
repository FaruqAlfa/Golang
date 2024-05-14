package main

import (
	"fmt"
	"net/http"
	"time"
)

func TimeHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		now := time.Now()
		hari := now.Weekday()
		day := now.Day()
		month := now.Month()
		year := now.Year()
		fmt.Println(day, month, year)
		result := fmt.Sprintf("%v, %v %v %v", hari, day, month, year)
		w.Write([]byte(result))
	}) // TODO: replace this
}

func SayHelloHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			w.Write([]byte("Hello there"))
		}else{
			w.Write([]byte("Hello, " + name + "!"))
		}
	}) // TODO: replace this
}

func main() {
	http.HandleFunc("/time", TimeHandler())
	http.HandleFunc("/hello", SayHelloHandler())	// TODO: answer here
	http.ListenAndServe("localhost:8080", nil)
}
