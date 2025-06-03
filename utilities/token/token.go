package utiltoken

import (
	"SimpleHTMLPage/config"
	"SimpleHTMLPage/consts"
	"SimpleHTMLPage/responses"
	"time"

	"github.com/golang-jwt/jwt"
)

type UserClaims struct {
	*jwt.StandardClaims
	UserRes *responses.UserResponse
}

func ParseUserToken(tokenStr string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &UserClaims{}, func(token *jwt.Token) (any, error) {
		return config.GetConfig().GetPublicKey(), nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, consts.ErrTokenInvalid
}

func CreateToken(userRes *responses.UserResponse) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("RS256"))

	// Define the token claims
	token.Claims = &UserClaims{
		&jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Second * time.Duration(config.GetConfig().GetJWTExpireDuration())).Unix(),
		},
		userRes,
	}

	// Private key is used here to signify that
	// The token is created by the server and only the server
	return token.SignedString(config.GetConfig().GetPrivateKey())
}
