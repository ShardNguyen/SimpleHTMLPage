package requests

import utilvalidate "SimpleHTMLPage/utilities/validate"

type UserLoginRequest struct {
	Username    string `json:"username"`
	RawPassword string `json:"password"`
}

type UserSignUpRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	RawPassword string `json:"password"`
}

func (loginReq *UserLoginRequest) CheckValidUsername() error {
	return utilvalidate.CheckValidUsername(loginReq.Username)
}

func (signUpReq *UserSignUpRequest) CheckValidInput() error {
	if err := utilvalidate.CheckValidUsername(signUpReq.Username); err != nil {
		return err
	}

	if err := utilvalidate.CheckValidEmail(signUpReq.Email); err != nil {
		return err
	}

	if err := utilvalidate.CheckValidPassword(signUpReq.RawPassword); err != nil {
		return err
	}

	return nil
}
