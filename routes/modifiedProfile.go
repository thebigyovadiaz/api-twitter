package routes

import (
	"encoding/json"
	"net/http"

	"github.com/thebigyovadiaz/api-twitter/db"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* ModifiedProfile is the func to modified the user register */
func ModifiedProfile(w http.ResponseWriter, r *http.Request) {
	var t *models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Incorrects fields "+err.Error(), http.StatusBadRequest)
		return
	}

	var status bool

	status, err = db.ModifiedRegister(t, UserID)
	if err != nil {
		http.Error(w, "Error to modified the user register "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Don't modified the user register "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
