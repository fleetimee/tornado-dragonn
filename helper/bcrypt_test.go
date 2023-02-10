package helper

import (
	"testing"
)

func TestHashPassword(t *testing.T) {
	password := "secretpassword"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Errorf("Error hashing password: %v", err)
	}
	if len(hashedPassword) == 0 {
		t.Errorf("Expected hashed password to have a value, got empty string")
	}
}

func TestCheckPasswordHash(t *testing.T) {
	password := "secretpassword"
	hashedPassword, _ := HashPassword(password)

	valid := CheckPasswordHash(password, hashedPassword)
	if !valid {
		t.Errorf("Expected valid hash, got invalid")
	}

	invalid := CheckPasswordHash("wrongpassword", hashedPassword)
	if invalid {
		t.Errorf("Expected invalid hash, got valid")
	}
}
