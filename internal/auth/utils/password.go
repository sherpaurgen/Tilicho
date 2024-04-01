package utils

import "golang.org/x/crypto/bcrypt"

func CreateSecretHash(secret string) (string, error) {

	salt, err := bcrypt.GenerateFromPassword([]byte(secret), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(salt), nil
}
