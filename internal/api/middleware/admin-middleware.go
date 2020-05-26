package middleware

import (
	"github.com/loremcookie/go-home/backend/internal/api/authentication"
	"net/http"
)

//AdminMiddleware is responsible for making a router only accessible for admin accounts.
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		//Get token from request
		token, err := authentication.GetTokenFromHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Extract metadata from token
		tokenMetadata := authentication.GetTokenMetadata(token)
		if tokenMetadata == nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//If admin isn't true return unauthorized status code
		if tokenMetadata.Admin != true {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Continue middleware chain
		next.ServeHTTP(w, r)
	})
}
