package auth

import "testing"

func TestHashAndVerifyPassword_Success(t *testing.T) {
	password := "secure-password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if err := VerifyPassword(hash, password); err != nil {
		t.Fatalf("expected password to be valid")
	}
}

func TestVerifyPassword_Failure(t *testing.T) {
	password := "secure-password"
	wrong := "wrong-password"

	hash, err := HashPassword(password)
	if err != nil {
		t.Fatalf("unexpected error hashing password: %v", err)
	}

	if err := VerifyPassword(hash, wrong); err == nil {
		t.Fatalf("expected invalid password error")
	}
}

func TestHashPassword_TooShort(t *testing.T) {
	_, err := HashPassword("short")
	if err == nil {
		t.Fatalf("expected error for short password")
	}
}
