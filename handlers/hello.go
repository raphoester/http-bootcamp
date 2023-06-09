package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	firstname := mux.Vars(r)["name"]
	w.Write([]byte(fmt.Sprintf("hello %s !", firstname)))

	UserDatabase = append(
		UserDatabase,
		User{
			Name: firstname,
		},
	)

	fmt.Printf("users list : %v\n", UserDatabase)
}
