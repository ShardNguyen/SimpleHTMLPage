package handlers

import (
	"SimpleHTMLPage/consts"
	"SimpleHTMLPage/models"
	"SimpleHTMLPage/requests"
	utilauth "SimpleHTMLPage/utilities/auth"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq requests.UserLoginRequest

	// Verify login request
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		RespondJSONError(w, http.StatusBadRequest, consts.InputInvalid)
		return
	}

	user, err := models.GetUser(&userReq)

	// Reconsidering about this? Since it is not recommended
	if err != nil {
		RespondJSONError(w, http.StatusUnauthorized, consts.UsernameInvalid)
		return
	}

	if !utilauth.VerifyPassword(user.Password, user.Salt, userReq.RawPassword) {
		RespondJSONError(w, http.StatusUnauthorized, consts.PasswordInvalid)
		return
	}

	token, err := utilauth.CreateToken(&userReq)

	if err != nil {
		RespondJSONError(w, http.StatusInternalServerError, consts.TokenGetFailed)
		return
	}

	RespondJSONOK(w, map[string]string{"token": token})
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userReq requests.UserSignUpRequest

	// Verify signup request
	err := json.NewDecoder(r.Body).Decode(&userReq)
	if err != nil {
		RespondJSONError(w, http.StatusBadRequest, consts.InputInvalid)
		return
	}

	err = models.CreateUser(&userReq)

	if err == consts.ErrUsernameExisted {
		RespondJSONError(w, http.StatusBadRequest, consts.UsernameExisted)
		return
	}

	if err != nil {
		RespondJSONError(w, http.StatusInternalServerError, consts.CreateFailed)
		return
	}

	RespondJSONOK(w, consts.Registered)
}

func (uh *UserHandler) Validate(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		RespondJSONError(w, http.StatusUnauthorized, "Missing authorization header")
		return
	}

	tokenString = tokenString[len("Bearer "):]
	fmt.Println(tokenString)
	userClaims, err := utilauth.ParseUserToken(tokenString)

	if err != nil {
		RespondJSONError(w, http.StatusUnauthorized, err.Error())
		return
	}

	RespondJSONOK(w, userClaims)
}

func (uh *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {

}
