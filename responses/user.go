package responses

import (
	"SimpleHTMLPage/models"
	"time"
)

type UserResponse struct {
	Username  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		Username:  user.Username,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}
