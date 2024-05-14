package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/route", func(w http.ResponseWriter, r *http.Request){
		// read request method
		method := r.Method

		//read URL and param value
		url := r.URL
		urlParam := r.URL.Query().Get("param")


		//read req header
		headerValues := r.Header.Values("foo") // menampilkan ["bar1", "bar2"]
		headerGet := r.Header.Get("foo")	   // menampilkan "bar1" . Hanya nilai pertama yang dikembalikan.

		//read req body
		body, err := io.ReadAll(r.Body)
		if err != nil{
		w.Write([]byte("error read body" + err.Error()))
		return
		}

		writer := fmt.Sprintf("Method:\t\t %v \nUrl:\t\t %v \nUrl param:\t %v \nHeaderValues:\t %v \nHeaderGet:\t %v \nBody:\t\t %v", method, url, urlParam, headerValues, headerGet, string(body))
		w.Write([]byte(writer))
	})

	fmt.Println("server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}