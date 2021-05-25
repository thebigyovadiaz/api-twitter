package routes

import (
	"log"
	"encoding/json"
	"net/http"

	"github.com/thebigyovadiaz/api-twitter/db"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* Register is the func to create an user in DB */
func Register(w http.ResponseWriter, r *http.Request) {
	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Bad Request, verify data sended"+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "Email user is required", http.StatusBadRequest)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "Password is too short", http.StatusBadRequest)
		return
	}

	_, encontrado, _ := db.CheckUserExist(t.Email)
	if encontrado == true {
		http.Error(w, "Email user exist, used other email", http.StatusBadRequest)
		return
	}

	_, status, err := db.InsertRegister(t)
	if err != nil {
		http.Error(w, "Error to insert in DB"+err.Error(), http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "Problem to insert in DB"+err.Error(), http.StatusBadRequest)
		return
	}

	log.Println("User Created!")
	w.WriteHeader(http.StatusCreated)
}
