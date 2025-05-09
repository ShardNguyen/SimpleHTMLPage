package handlers

import (
	"net/http"
)

type IAuthHandler interface {
	Login(w http.ResponseWriter, r *http.Request)
	SignUp(w http.ResponseWriter, r *http.Request)
	SignOut(w http.ResponseWriter, r *http.Request)
}

type IOrderHandler interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
	ListOrders()
	GetOrder()
}
