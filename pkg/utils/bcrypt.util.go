package utils

import "golang.org/x/crypto/bcrypt"

/*
*
HashPassword

	func HashPassword oneway hash only password
*/
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

/*
*
CheckPasswordHash

	func CheckPasswordHash check hash
*/
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
