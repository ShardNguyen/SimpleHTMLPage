package responses

import "fmt"

type UserResponse struct {
	Username string `json:"username"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password"`
}

func (userRes *UserResponse) PrintInfo() {
	fmt.Println("Username: ", userRes.Username)
	fmt.Println("Email: ", userRes.Email)
	fmt.Println("Password: ", userRes.Password)
}
