package helpers

import (
	"crypto/ed25519"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"github.com/pkg/errors"
	"golang.org/x/crypto/scrypt"
)

type SignatureData struct {
	UserId  int64 `json:"user_id"`
	OfferId int64 `json:"offer_id"`
}

func GenerateKeys(email, password, salt []byte, n, r, p int) (string, string, error) {
	emailSalt := append(salt, email...)
	emailSaltHashed := sha256.Sum256(emailSalt)

	cypherKey, err := scrypt.Key(password, emailSaltHashed[:], n, r, p, ed25519.SeedSize)
	if err != nil {
		return "", "", err
	}

	accountSK := ed25519.NewKeyFromSeed(cypherKey)

	return hex.EncodeToString(accountSK[:]), hex.EncodeToString(accountSK.Public().(ed25519.PublicKey)[:]), nil
}

func Sign(sk string, data SignatureData) (string, error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return "", errors.Wrap(err, "failed to marshal data")
	}
	messageHashed := sha256.Sum256(marshaled)
	decodedSk, _ := hex.DecodeString(sk)
	return hex.EncodeToString(ed25519.Sign(decodedSk, messageHashed[:])), nil
}

func Verify(pk string, message []byte, sig string) (bool, error) {
	messageHashed := sha256.Sum256(message)
	decodedSk, err := hex.DecodeString(pk)
	if err != nil {
		return false, err
	}
	decodedSig, err := hex.DecodeString(sig)
	if err != nil {
		return false, err
	}

	return ed25519.Verify(decodedSk, messageHashed[:], decodedSig), nil
}
