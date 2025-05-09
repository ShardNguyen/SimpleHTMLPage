package requests

type UserLoginRequest struct {
	Username    string `json:"username"`
	RawPassword string `json:"password"`
}

type UserSignUpRequest struct {
	Username    string `json:"username"`
	Email       string `json:"email"`
	RawPassword string `json:"password"`
}

func (signUpReq UserSignUpRequest) ConvertToUserLoginRequest() *UserLoginRequest {
	return &UserLoginRequest{
		Username:    signUpReq.Username,
		RawPassword: signUpReq.RawPassword,
	}
}
