package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"
	"log"
)

func encrypt(plaintext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {

		log.Println(err)
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

func decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}

// Encrypt encrypts a plaintext using a pre-defined key
func Encrypt(plaintext, key string) (string, error) {

	var ciphertext []byte
	var err error

	ciphertext, err = encrypt([]byte(plaintext), []byte(key))

	return base64.StdEncoding.EncodeToString(ciphertext), err
}

// Decrypt decrypts a ciphertext given a pre-defined 32 bit key
func Decrypt(ciphertext, key string) (string, error) {
	var plaintext, ciphertextb []byte
	var err error

	ciphertextb, err = base64.StdEncoding.DecodeString(string(ciphertext))
	plaintext, err = decrypt(ciphertextb, []byte(key))

	return string(plaintext), err
}
