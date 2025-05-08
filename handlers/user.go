package handlers

import (
	"SimpleHTMLPage/models"
	"SimpleHTMLPage/responses"
	"SimpleHTMLPage/utilities"
	"encoding/json"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userRes responses.UserResponse

	if err := json.NewDecoder(r.Body).Decode(&userRes); err != nil {
		utilities.RespondError(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	user, err := models.GetUser(&userRes)

	if err != nil {
		utilities.RespondError(w, http.StatusInternalServerError, "Failed to get username")
		return
	}

	if user.ID == 0 {
		utilities.RespondError(w, http.StatusUnauthorized, "Invalid username")
		return
	}

	if !utilities.VerifyPassword(user.Password, user.Salt, userRes.Password) {
		utilities.RespondError(w, http.StatusUnauthorized, "Invalid password")
		return
	}

	utilities.RespondOK(w, "Logged in")
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userRes responses.UserResponse

	if err := json.NewDecoder(r.Body).Decode(&userRes); err != nil {
		utilities.RespondError(w, http.StatusBadRequest, "Invalid Input")
		return
	}

	if err := models.CreateUser(&userRes); err != nil {
		utilities.RespondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	utilities.RespondOK(w, "Registered")
}

func (uh *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {

}
