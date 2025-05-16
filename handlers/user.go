package handlers

import (
	"encoding/json"
	"net/http"

	"SimpleHTMLPage/consts"
	dbtoken "SimpleHTMLPage/databases/token"
	"SimpleHTMLPage/models"
	"SimpleHTMLPage/requests"
	"SimpleHTMLPage/responses"
	utilpass "SimpleHTMLPage/utilities/password"
	utiltoken "SimpleHTMLPage/utilities/token"
)

type UserHandler struct {
}

func NewUserHandler() *UserHandler {
	return &UserHandler{}
}

func (uh *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var userReq requests.UserLoginRequest

	// Check if requested json file is decodable
	// Also check input validity
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil || !userReq.CheckValidInput() {
		RespondJSONError(w, http.StatusBadRequest, consts.InputInvalid)
		return
	}

	user, err := models.GetUser(userReq.Username)

	// Reconsidering about this? Since it is not recommended
	if err != nil {
		RespondJSONError(w, http.StatusUnauthorized, consts.UsernameInvalid)
		return
	}

	if !utilpass.VerifyPassword(user.Password, user.Salt, userReq.RawPassword) {
		RespondJSONError(w, http.StatusUnauthorized, consts.PasswordInvalid)
		return
	}

	userRes := responses.NewUserResponse(user)
	token, err := utiltoken.CreateToken(userRes)

	if err != nil {
		RespondJSONError(w, http.StatusInternalServerError, consts.TokenGetFailed)
		return
	}

	dbtoken.AddToken(token)
	dbtoken.PrintAmountOfTokens()
	RespondJSONOK(w, map[string]string{"token": token})
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userReq requests.UserSignUpRequest

	// Check if requested json file is decodable
	// Also check input validity
	err := json.NewDecoder(r.Body).Decode(&userReq)

	if err != nil || !userReq.CheckValidInput() {
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

	RespondJSONOK(w, map[string]string{"message": consts.Registered})
}

func (uh *UserHandler) ValidateUser(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		RespondJSONError(w, http.StatusUnauthorized, consts.AuthHeaderMissing)
		return
	}

	tokenString = tokenString[len("Bearer "):]
	userClaims, err := utiltoken.ParseUserToken(tokenString)

	// Check validity of the request token
	if err != nil || !dbtoken.CheckUserTokenExists(tokenString) {
		RespondJSONError(w, http.StatusUnauthorized, consts.TokenInvalid)
		dbtoken.DeleteToken(tokenString) // In case token is expired but is still inside the token db
		return
	}

	userResponse := userClaims.UserRes
	RespondJSONOK(w, map[string]responses.UserResponse{"data": *userResponse})
}

func (uh *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")

	if tokenString == "" {
		RespondJSONError(w, http.StatusUnauthorized, consts.AuthHeaderMissing)
		return
	}

	tokenString = tokenString[len("Bearer "):]
	dbtoken.DeleteToken(tokenString)
	dbtoken.PrintAmountOfTokens()
}
