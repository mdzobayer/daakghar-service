package hash

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// GenerateHashFromPassword generate hash
func GenerateHashFromPassword(password string) (hash string, err error) {

	if hash, err = generateHashFromPassword(password); err != nil {
		log.Println("hash : GenerateHashFromPassword : ", err)
	}

	return
}

// CompareHashAndPassword compare hashFromDb and userPassword
func CompareHashAndPassword(hashFromDb string, userPassword string) (err error) {

	if err = compareHashAndPassword(hashFromDb, userPassword); err != nil {
		log.Println("hash: CompareHashAndPassword : ", err)
	}

	return
}

func generateHashFromPassword(password string) (hash string, err error) {
	var hashByte []byte

	if hashByte, err = bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost); err != nil {
		log.Println("hash: generateHashFromPassword : ", err)
	}

	hash = string(hashByte[:])

	return
}

func compareHashAndPassword(hashFromDb, userPassword string) (err error) {
	if err = bcrypt.CompareHashAndPassword([]byte(hashFromDb), []byte(userPassword)); err != nil {
		log.Println("hash : compareHashAndPassword : ", err)
	}

	return
}
