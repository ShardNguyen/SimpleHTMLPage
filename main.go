package main

import (
	"SimpleHTMLPage/config"
	dbpostgres "SimpleHTMLPage/databases/postgresql"
	"SimpleHTMLPage/handlers"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	err := config.ParseConfig()

	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()

	err = dbpostgres.UserConnect()

	if err != nil {
		fmt.Println("Cannot connect to database")
		return
	}

	userHandler := handlers.NewUserHandler()

	r.HandleFunc("/signup", userHandler.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/login", userHandler.Login).Methods(http.MethodPost)

	http.ListenAndServe(":8080", r)
}
