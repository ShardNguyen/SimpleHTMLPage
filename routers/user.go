package routers

import (
	"SimpleHTMLPage/handlers"
	"SimpleHTMLPage/middlewares"
	"net/http"
)

var userHandler = handlers.NewUserHandler()

func userRouter() {
	r.HandleFunc("/user/signup", userHandler.SignUp).Methods(http.MethodPost)
	r.HandleFunc("/user/login", userHandler.Login).Methods(http.MethodPost)
	// Before getting user's info, go through the auth middleware
	r.Handle("/user/info", middlewares.Authenticate(http.HandlerFunc(userHandler.ShowCurrentUserInfo))).Methods(http.MethodGet)
	// Before signout is reached, it has to through the auth middleware
	r.Handle("/user/signout", middlewares.Authenticate(http.HandlerFunc(userHandler.SignOut))).Methods(http.MethodPost)
}
