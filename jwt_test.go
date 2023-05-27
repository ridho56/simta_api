package main

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
	"goelster/helper"
	"golang.org/x/crypto/bcrypt"
	"testing"
	"time"
)

func TestJWTValidate(t *testing.T) {
	hashedPasswordFromDatabase := "$2y$10$TxXrUsMRkbOqtnmbTtBt0OLqk1HUBdzHUDSwSepJa3ysiNLJTsM5S"
	inputPasswordFromUser := "Haha123_"
	if helper.CompareHashAndPasswordSha384(hashedPasswordFromDatabase, inputPasswordFromUser) {
		fmt.Println("Password matches!")
	} else {
		fmt.Println("Password does not match!")
	}
}

func TestCreateJWT(t *testing.T) {
	jwt := helper.CreateNewJWT("Rendra", "ridho@gmail.com", "secret-key", time.Hour*5)
	fmt.Println(jwt)
}

func TestCompareHash(t *testing.T) {
	password := "11102002"
	bcydatabase := []byte("$2y$10$dm1CVCjzHMqmbV8iFWjQueJsvOA6RceyCWBr4UMfh2VI9ND7kFKW6")

	// Hash using SHA-384 and encode as base64
	passwordBytes := []byte(password)
	hasher := sha512.New384()
	hasher.Write(passwordBytes)
	sha384Hash := base64.StdEncoding.EncodeToString(hasher.Sum(nil))
	preparedPassword := []byte(sha384Hash)

	// Hash using Bcrypt
	cost := 10
	bcryptHash, err := bcrypt.GenerateFromPassword(preparedPassword, cost)
	if err != nil {
		fmt.Println("Error generating hash:", err)
		return
	}

	fmt.Println("SHA-384 hash:", sha384Hash)
	fmt.Println("Bcrypt hash:", string(bcryptHash))

	// Compare password
	err = bcrypt.CompareHashAndPassword(bcydatabase, []byte(sha384Hash))
	if err == bcrypt.ErrMismatchedHashAndPassword {
		fmt.Println("Password does not match!")
	} else if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Password matches!")
	}
}
