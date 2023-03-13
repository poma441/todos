package hasher

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

/*
*	Функция генерации рандомного хэша SHA-256
 */
func GenerateRandomHash256() (string, error) {
	randomBytes := make([]byte, 32)

	if _, err := rand.Read(randomBytes); err != nil {
		return "", err
	}

	hash := sha256.Sum256(randomBytes)
	rezSha := hex.EncodeToString(hash[:])

	return rezSha, nil
}

/*
*	Функция генерации хэша пароля
 */
func HashPass(inputPass string) string {
	hashPass, err := bcrypt.GenerateFromPassword([]byte(inputPass), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(hashPass)
}

/*
*	Функция сравнения хэша пароля из базы и пароля в чистом виде, предоставленного пользователем
 */
func ComparePass(userHashPass string, inputPass string) error {
	err := bcrypt.CompareHashAndPassword([]byte(userHashPass), []byte(inputPass))
	if err != nil {
		return errors.New("неверный пароль")
	}

	return err
}
