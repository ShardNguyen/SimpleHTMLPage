package utilauth

import (
	"SimpleHTMLPage/config"
	"SimpleHTMLPage/consts"
	"SimpleHTMLPage/requests"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	*jwt.StandardClaims
	UserReq *requests.UserLoginRequest
}

func ParseUserToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(config.GetConfig().GetSecretKey()), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, consts.ErrTokenInvalid
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
