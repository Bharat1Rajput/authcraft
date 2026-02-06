package auth

import (
	"testing"
	"time"
)

func TestGenerateAndVerifyAccessToken(t *testing.T) {
	secret := []byte("test-secret")
	tm := NewTokenManager(secret, time.Minute)

	userID := "user-123"
	role := "admin"

	token, err := tm.GenerateAccessToken(userID, role)
	if err != nil {
		t.Fatalf("failed to generate token: %v", err)
	}

	gotUserID, gotRole, err := tm.VerifyAccessToken(token)
	if err != nil {
		t.Fatalf("failed to verify token: %v", err)
	}

	if gotUserID != userID || gotRole != role {
		t.Fatalf("unexpected claims")
	}
}

func TestVerifyAccessToken_InvalidToken(t *testing.T) {
	secret := []byte("test-secret")
	tm := NewTokenManager(secret, time.Minute)

	if _, _, err := tm.VerifyAccessToken("invalid.token"); err == nil {
		t.Fatalf("expected invalid token error")
	}
}
