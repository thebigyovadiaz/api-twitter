package middlewares

import (
	"net/http"

	"github.com/thebigyovadiaz/api-twitter/routes"
)

/* ValidateJWT is the func to validate the JWT in the request */
func ValidateJWT(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, _, _, err := routes.ProcessToken(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "Token Error: "+err.Error(), http.StatusBadRequest)
			return
		}

		next.ServeHTTP(w, r)
	}
}
