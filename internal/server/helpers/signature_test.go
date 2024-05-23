package helpers

import (
	"crypto/ed25519"
	"encoding/hex"
	"encoding/json"
	"testing"
)

func TestGenerateKeys(t *testing.T) {
	email := []byte("test@example.com")
	password := []byte("password123")
	salt := []byte("random_salt")
	n, r, p := 16384, 8, 1

	sk, pk, err := GenerateKeys(email, password, salt, n, r, p)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(sk) != hex.EncodedLen(ed25519.PrivateKeySize) {
		t.Errorf("Expected secret key of length %d, got %d", hex.EncodedLen(ed25519.PrivateKeySize), len(sk))
	}

	if len(pk) != hex.EncodedLen(ed25519.PublicKeySize) {
		t.Errorf("Expected public key of length %d, got %d", hex.EncodedLen(ed25519.PublicKeySize), len(pk))
	}
}

func TestSign(t *testing.T) {
	data := SignatureData{
		UserId:  12345,
		OfferId: 67890,
	}

	email := []byte("test@example.com")
	password := []byte("password123")
	salt := []byte("random_salt")
	n, r, p := 16384, 8, 1

	sk, _, err := GenerateKeys(email, password, salt, n, r, p)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	signature, err := Sign(sk, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(signature) != hex.EncodedLen(ed25519.SignatureSize) {
		t.Errorf("Expected signature of length %d, got %d", hex.EncodedLen(ed25519.SignatureSize), len(signature))
	}
}

func TestVerify(t *testing.T) {
	data := SignatureData{
		UserId:  12345,
		OfferId: 67890,
	}

	email := []byte("test@example.com")
	password := []byte("password123")
	salt := []byte("random_salt")
	n, r, p := 16384, 8, 1

	sk, pk, err := GenerateKeys(email, password, salt, n, r, p)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	signature, err := Sign(sk, data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	marshaledData, err := json.Marshal(data)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	valid, err := Verify(pk, marshaledData, signature)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if !valid {
		t.Errorf("Expected signature to be valid")
	}

	// Test with an invalid signature
	invalidSig := signature[:len(signature)-1] + "0"
	valid, err = Verify(pk, marshaledData, invalidSig)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if valid {
		t.Errorf("Expected signature to be invalid")
	}

	// Test with modified data
	modifiedData := SignatureData{
		UserId:  54321,
		OfferId: 67890,
	}

	marshaledModifiedData, err := json.Marshal(modifiedData)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	valid, err = Verify(pk, marshaledModifiedData, signature)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if valid {
		t.Errorf("Expected signature to be invalid for modified data")
	}
}
