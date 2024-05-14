package server

import (
	"encoding/json"
	"net/http"

	"github.com/rest/model"
)

func Students(w http.ResponseWriter, r *http.Request){
	data := []model.Student{
		{Name : "Fahmi", Age : 22},
		{Name : "Rizki", Age : 23},
		{Name : "Aldi", Age : 24},
		{Name : "Akbar", Age : 25},
	}

	jsonInBytes, err := json.Marshal(data)
	if err != nil{
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonInBytes)

}