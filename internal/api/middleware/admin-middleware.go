package middleware

import (
	"github.com/loremcookie/go-home/backend/internal/api/authentication"
	"github.com/loremcookie/go-home/backend/internal/listutill"
	"net/http"
	"os"
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

		//Get list of routes to exclude from admin privileges
		excludeRoutes := listutill.ParseList(os.Getenv("ADMIN_EXCLUDE_ROUTES"))
		//Get request path
		requestPath := r.URL.Path
		//Compare request path to exclude paths
		for _, route := range excludeRoutes {
			if route == requestPath {

				//When route matches a eluded route
				//then validate the token
				err = authentication.VerifyToken(token)
				if err != nil {
					w.WriteHeader(http.StatusForbidden)
					return
				}

				//When the token is valid the continue the middleware chain
				next.ServeHTTP(w, r)
				return
			}
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
