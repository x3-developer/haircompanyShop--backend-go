package valueobject

import (
	"golang.org/x/crypto/bcrypt"
)

type PasswordHash string

func NewPasswordHash(raw string) (PasswordHash, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(raw), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return PasswordHash(hash), nil
}

func (p PasswordHash) Check(raw string) bool {
	return bcrypt.CompareHashAndPassword([]byte(p), []byte(raw)) == nil
}
