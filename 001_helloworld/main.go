package main

import (
	"fmt"
	"net/http"
)

func helloworld_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "Hello World!")
}
func helloname_handler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello %s!", r.PathValue("name"))
}

func main(){
	mux:=http.NewServeMux()
	mux.HandleFunc("GET /",helloworld_handler)
	mux.HandleFunc("GET /{name}",helloname_handler)
	http.ListenAndServe(":8080",mux)
	
}

/*
GET -> SELECT
POST -> INSERT
DELETE -> DELETE
*/

//go run main.go
//go build -o main.go