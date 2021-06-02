package routes

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/thebigyovadiaz/api-twitter/db"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* InsertTweet: save the tweet in DB */
func InsertTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		http.Error(w, "Message is required "+err.Error(), http.StatusBadRequest)
		return
	}

	register := models.InsertTweet{
		UserID:    UserID,
		Message:   message.Message,
		CreatedAt: time.Now(),
	}

	_, status, error := db.InsertTweet(register)
	if error != nil {
		http.Error(w, "Error to insert register, retry "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "Don't insert register "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
