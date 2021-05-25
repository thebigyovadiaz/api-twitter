package middlewares

import (
	"net/http"
	"github.com/thebigyovadiaz/api-twitter/db"
)

/* CheckDB is the middleware to check connect DB */
func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if db.CheckingConnection() == 0 {
			http.Error(w, "Connect loss to DB", 500)
			return
		}

		next.ServeHTTP(w, r)
	}
}
