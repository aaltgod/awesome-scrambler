package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	cryptoRand "crypto/rand"
	"encoding/base64"
	"io"
	"log"
	"math/rand"
	"time"
)

func Encrypt(plainTextString string) (string, string, error) {

	rand.Seed(time.Now().UnixNano())

	keyBytes := make([]byte, 32)
	if _, err := rand.Read(keyBytes); err != nil {
		log.Println("[RAND-READ]: ", err)
		return "", "", err
	}


	key := base64.StdEncoding.EncodeToString(keyBytes)
	plaintext := []byte(plainTextString)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		log.Println("[NEW-CIPHER]: ", err)
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("[NEW-GCM: ", err)
		return "", "", err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(cryptoRand.Reader, nonce); err != nil {
		log.Println("[READ-FULL]: ", err)
		return "", "", err
	}

	cipherTextBytes := aesGCM.Seal(nonce, nonce, plaintext, nil)
	cipherText := base64.StdEncoding.EncodeToString(cipherTextBytes)

	return cipherText, key, nil
}