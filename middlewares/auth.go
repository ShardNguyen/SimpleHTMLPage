// Middlewares take a request and do something with it
// Then it passes the request down to another middleware or the final handler

package middlewares

import (
	"SimpleHTMLPage/consts"
	dbredis "SimpleHTMLPage/databases/redis"
	utilresponders "SimpleHTMLPage/utilities/responders"
	utiltoken "SimpleHTMLPage/utilities/token"
	"fmt"
	"net/http"
)

func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.AuthHeaderMissing)
			return
		}

		tokenString = tokenString[len("Bearer "):]

		_, err := dbredis.GetTokenStorage().Get(r.Context(), tokenString).Result()

		if err != nil {
			utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.TokenNotExists)
			return
		}

		userClaims, err := utiltoken.ParseUserToken(tokenString)

		if err != nil {
			utilresponders.RespondJSONError(w, http.StatusBadRequest, consts.TokenGetFailed)
			return
		}

		userResponse := userClaims.UserRes
		fmt.Println(userResponse)

		// If return is not reached, proceed to the "next" handler
		next.ServeHTTP(w, r)
	})
}
