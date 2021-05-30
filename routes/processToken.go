package routes

import (
	"log"
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thebigyovadiaz/api-twitter/db"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* Email is the email token user valid */
var Email string

/* UserID is the id token user valid */
var UserID string

/* ProcessToken is the func process and validate user token */
func ProcessToken(token string) (*models.Claim, bool, string, error) {
	myPass := []byte("MastersDelDesarrollo_grupoDeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("format token invalid")
	}

	token = strings.TrimSpace(splitToken[1])

	tokn, err := jwt.ParseWithClaims(token, claims, func(tk *jwt.Token) (interface{}, error) {
		return myPass, nil
	})

	if err != nil {
		return claims, false, string(""), err
	}

	if !tokn.Valid {
		return claims, false, string(""), errors.New("invalid token")
	}

	user, searched, _ := db.CheckUserExist(claims.Email)
	if searched {
		Email = user.Email
		UserID = user.ID.Hex()
	}

	return claims, searched, UserID, nil
}

