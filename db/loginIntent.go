package db

import (
	"github.com/thebigyovadiaz/api-twitter/models"
	"golang.org/x/crypto/bcrypt"
)

/* LoginIntent is the func to check login's DB */
func LoginIntent(email string, password string) (models.User, bool) {
	user, exist, _ := CheckUserExist(email)
	if !exist {
		return user, exist
	}

	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)
	
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}

	return user, true
}
