package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thebigyovadiaz/api-twitter/db"
	"github.com/thebigyovadiaz/api-twitter/jwt"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* Login is the func to login APP */
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "User - Pasword invalid "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "User email in required", http.StatusBadRequest)
		return
	}

	user, exist := db.LoginIntent(t.Email, t.Password)
	if !exist {
		http.Error(w, "Ocurried and error to login", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(user)
	if err != nil {
		http.Error(w, "Error to generate the token "+err.Error(), http.StatusBadRequest)
		return
	}

	response := models.LoginResponse{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(response)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "toke",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
