package db

import "golang.org/x/crypto/bcrypt"

/* PasswordEncrypt is a func for encrypt password */
func PasswordEncrypt(pass string) (string, error) {
	costo := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), costo)
	return string(bytes), err
}
