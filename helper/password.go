package helper

import (
	"crypto/sha512"
	"encoding/base64"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	PanicIfError(err)

	return true, nil
}

func CompareHashAndPasswordSha384(hash string, password string) bool {
	passwordBytes := []byte(password)
	hashes := sha512.New384()
	hashes.Write(passwordBytes)
	sha384Hash := base64.StdEncoding.EncodeToString(hashes.Sum(nil))
	preparedPassword := []byte(sha384Hash)

	err := bcrypt.CompareHashAndPassword([]byte(hash), preparedPassword)
	return err == nil
}
