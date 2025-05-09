package utilauth

import (
	"SimpleHTMLPage/config"
	"SimpleHTMLPage/requests"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	*jwt.StandardClaims
	UserReq *requests.UserLoginRequest
}

func CreateToken(userReq *requests.UserLoginRequest) (string, error) {
	var secretKey = config.GetConfig().GetSecretKey()

	// Define the token claims
	claims := &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
		},
		userReq,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(secretKey))
}
