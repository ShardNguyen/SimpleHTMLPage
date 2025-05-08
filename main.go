package main

import (
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	"SimpleHTMLPage/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	err := dbpostgres.UserConnect()

	if err != nil {
		fmt.Println("Cannot connect to database")
		return
	}

	uh := handlers.NewUserHandler()
	r.HandleFunc("POST /signup", uh.SignUp)
	r.HandleFunc("POST /login", uh.Login)

	http.ListenAndServe(":8080", nil)
}
