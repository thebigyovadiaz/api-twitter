package jwt

import (
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/thebigyovadiaz/api-twitter/models"
)

/* GenerateJWT is the func to generate a token with JWT */
func GenerateJWT(t models.User) (string, error) {
	myPass := []byte("MastersDelDesarrollo_grupoDeFacebook")

	payload := jwt.MapClaims{
		"email":     t.Email,
		"firstname": t.Firstname,
		"lastname":  t.Lastname,
		"birthday":  t.Birthday,
		"biography": t.Biography,
		"location":  t.Location,
		"website":   t.Website,
		"_it":       t.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myPass)
	if err != nil {
		return tokenStr, err
	}

	log.Printf("Token Generated: %s\n", tokenStr)
	return tokenStr, nil
}
