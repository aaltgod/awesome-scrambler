package encrypter

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"math/rand"
)

func Encrypt(plainTextString string) (string, string, error) {

	bytes := make([]byte, 32)
	if _, err := rand.Read(bytes); err != nil {
		log.Println("[RAND-READ]: ", err)
		return "", "", err
	}

	keyString := hex.EncodeToString(bytes)
	key, _ := hex.DecodeString(keyString)
	plaintext := []byte(plainTextString)
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Println("[NEW-CIPHER]: ", err)
		return "", "", err
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Println("[NEW-GCM: ", err)
		return "", "", err
	}

	nonce := make([]byte, aesGCM.NonceSize())
	cipherText := aesGCM.Seal(nonce, nonce, plaintext, nil)

	return fmt.Sprintf("%x", cipherText), keyString, nil
}