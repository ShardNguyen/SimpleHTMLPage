package handlers

import (
	"net/http"
)

type Handler interface {
	Login(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	SignOut(w http.ResponseWriter, r *http.Request)
}
