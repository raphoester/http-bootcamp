package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/raphoester/http-bootcamp/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello/{name}", handlers.Hello).Methods("GET")
	fmt.Println("Starting server on port 8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		fmt.Printf("Failed starting server | %s", err.Error())
	}
}
