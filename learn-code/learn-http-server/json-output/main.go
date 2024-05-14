package main

import "fmt"
import "net/http"
import "encoding/json"

type student struct {
    Name string `json:"name"`
    Age  int    `json:"age"`
}

func ActionIndex(w http.ResponseWriter, r *http.Request) {
    data := []student{
        {"Aditira", 22},
        {"Dito", 24},
        {"Ojan", 30},
        {"Tegar", 25},
		{"Alfin", 23},
    }

    // jsonInBytes, err := json.Marshal(data)
    // if err != nil {
    //     http.Error(w, err.Error(), http.StatusInternalServerError)
    //     return
    // }

    // w.Header().Set("Content-Type", "application/json")
    // w.WriteHeader(http.StatusOK)
    // w.Write(jsonInBytes)

		//kode di atas ekuivalen dengan kode di bawah

	w.Header().Set("Content-Type", "application/json")

		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
}

func main() {
    http.HandleFunc("/", ActionIndex)

    fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", nil)
}