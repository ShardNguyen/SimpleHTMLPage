package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"SimpleHTMLPage/consts"
	dbredis "SimpleHTMLPage/databases/redis"
	"SimpleHTMLPage/models"
	"SimpleHTMLPage/requests"
	"SimpleHTMLPage/responses"
	utilpass "SimpleHTMLPage/utilities/password"
	utilresponders "SimpleHTMLPage/utilities/responders"
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
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.InputInvalid)
		return
	}

	// Check username's validity
	if err := userReq.CheckValidUsername(); err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.UsernameInvalid)
		return
	}

	// Getting user's info based on username
	user, err := models.GetUser(userReq.Username)

	// Check username
	// Reconsidering about this? Since it is not recommended
	if err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusUnauthorized, consts.UsernameInvalid)
		return
	}

	// Check password
	if !utilpass.VerifyPassword(user.Password, user.Salt, userReq.RawPassword) {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusUnauthorized, consts.PasswordInvalid)
		return
	}

	// Creating a user response to add to the token
	userRes := responses.NewUserResponse(user)
	userBytes, err := json.Marshal(userRes)
	if err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusInternalServerError, consts.JSONMarshalFailed)
		return
	}

	// Create a token with user response's info
	token, err := utiltoken.CreateToken(userRes)

	if err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusInternalServerError, consts.TokenGetFailed)
		return
	}

	// Store the created token into the redis database
	err = dbredis.StoreToken(token, userBytes)

	if err != nil {
		fmt.Println(err)
		utilresponders.RespondJSONError(w, http.StatusInternalServerError, consts.TokenStoreFailed)
		return
	}

	// Returns a 200 OK response with token as an additional message
	utilresponders.RespondJSONOK(w, map[string]string{"token": token})
}

func (uh *UserHandler) SignUp(w http.ResponseWriter, r *http.Request) {
	var userReq requests.UserSignUpRequest

	// Check if requested json file is decodable
	if err := json.NewDecoder(r.Body).Decode(&userReq); err != nil {
		utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.InputInvalid)
		return
	}

	// Check input
	if err := userReq.CheckValidInput(); err != nil {
		utilresponders.RespondJSONError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Create a new user based on user request's info
	err := models.CreateUser(&userReq)

	// Check if username existed
	if err == consts.ErrUsernameExisted {
		utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.UsernameExisted)
		return
	}

	if err != nil {
		utilresponders.RespondJSONError(w, http.StatusInternalServerError, consts.CreateFailed)
		return
	}

	// Returns a 200 OK response
	utilresponders.RespondJSONOK(w, map[string]string{"message": consts.SignedUp})
}

func (uh *UserHandler) ShowCurrentUserInfo(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	tokenString = tokenString[len("Bearer "):]

	data, err := dbredis.GetDataFromToken(tokenString)
	if err != nil {
		utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.TokenGetFailed)
		return
	}

	utilresponders.RespondJSONOK(w, map[string]any{"data": data})
}

func (uh *UserHandler) SignOut(w http.ResponseWriter, r *http.Request) {
	// Get token from the Authorization header
	tokenString := r.Header.Get("Authorization")

	// Check if token is empty
	if tokenString == "" {
		utilresponders.RespondJSONError(w, http.StatusUnauthorized, consts.AuthHeaderMissing)
		return
	}

	// Get the token string
	tokenString = tokenString[len("Bearer "):]

	// Delete token in redis database
	dbredis.DeleteToken(tokenString)

	// Returns a 200 OK response
	utilresponders.RespondJSONOK(w, map[string]string{"data": consts.SignedOut})
}
