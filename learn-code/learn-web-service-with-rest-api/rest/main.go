package main

import (
	"fmt"
	"net/http"

	"github.com/rest/client"
	"github.com/rest/server"
)

func main(){
	mux := http.NewServeMux()

	//server
	mux.HandleFunc("/api/student", server.Students)

	//Client
	mux.HandleFunc("/page/student", client.Student)

	fmt.Println("start server")
	http.ListenAndServe(":8080", mux)
}

