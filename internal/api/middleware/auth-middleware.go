package middleware

import (
	"github.com/loremcookie/go-home/backend/internal/api/authentication"
	"net/http"
	"os"
	"strings"
)

//AuthenticationMiddleware is the middleware hat validates the jwt and returns an unauthorized http status.
func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var err error

		//List of routes that are excluded by the auth middleware
		excludeAuth := strings.Split(strings.TrimSpace(os.Getenv("AUTH_EXCLUDE_ROUTES")), ",")
		//Trim spaces from list
		for id, value := range excludeAuth {
			excludeAuth[id] = strings.Trim(value, " ")
		}
		//Get path of request
		requestPath := r.URL.Path

		//Compare request path to list of paths to exclude from middleware
		for _, v := range excludeAuth {
			if v == requestPath {
				next.ServeHTTP(w, r)
				return
			}
		}

		//Get token from request
		token, err := authentication.GetTokenFromHeader(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		//Verify token
		err = authentication.VerifyToken(token)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
