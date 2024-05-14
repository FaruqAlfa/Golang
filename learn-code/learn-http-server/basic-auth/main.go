package main

import "net/http"

const USERNAME = "fahmi"
const PASSWORD = "12345"

func login(w http.ResponseWriter, r *http.Request){
	username, password, ok := r.BasicAuth()
	if !ok {
		w.Write([]byte("something went wrong"))
		return
	}
	isValid := (username == USERNAME) && (password == PASSWORD)
	if !isValid {
		w.Write([]byte("login failed"))
		return
	}
	w.Write([]byte ("login success"))
}

func main(){
http.HandleFunc("/login", login)

server := new(http.Server)
    server.Addr = ":8080"
    server.ListenAndServe()
}
