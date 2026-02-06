package security

import (
	"errors"
	"os"
)

var ErrMissingJWTSecret = errors.New("missing JWT secret")

func LoadJWTSecret() ([]byte, error) {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		return nil, ErrMissingJWTSecret
	}
	return []byte(secret), nil
}
