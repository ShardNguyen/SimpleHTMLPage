package requests

import utilinpvalid "SimpleHTMLPage/utilities/inpvalid"

type UserLoginRequest struct {
	Username    string `json:"username"`
	RawPassword string `json:"password"`
}

type UserSignUpRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	RawPassword string `json:"password"`
}

func (loginReq *UserLoginRequest) CheckValidInput() bool {
	return utilinpvalid.CheckValidUsername(loginReq.Username) &&
		utilinpvalid.CheckValidPassword(loginReq.RawPassword)
}

func (signUpReq *UserSignUpRequest) CheckValidInput() bool {
	return utilinpvalid.CheckValidUsername(signUpReq.Username) &&
		utilinpvalid.CheckValidEmail(signUpReq.Email) &&
		utilinpvalid.CheckValidPassword(signUpReq.RawPassword)
}
