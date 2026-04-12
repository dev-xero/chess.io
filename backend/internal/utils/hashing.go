package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/argon2"
)

// Argon2 customizable parameters.
type Argon2Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}

// Generates an encoded password hashed with salting.
//
// This external function encodes the generated hash and salt into a
// format safe enough for storage.
func GenerateEncodedPasswordHash(password string, p *Argon2Params) (string, error) {
	hash, salt, err := generatePasswordHash(password, p)
	if err != nil {
		return "", err
	}
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)
	return fmt.Sprintf(
		"$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash,
	), nil
}

// Generates a cryptographically secure password hash using Argon2id.
//
// This function receives a "hashing" configuration via the `Argon2Params`
// struct then generates a new Argon2ID key from the password string.
//
// This hash is used in conjunction with the salt hash for encoding.
func generatePasswordHash(password string, p *Argon2Params) (hash []byte, salt []byte, err error) {
	salt, err = generateSaltFromRandomBytes(p.SaltLength)
	if err != nil {
		return nil, nil, err
	}
	hash = argon2.IDKey(
		[]byte(password),
		salt,
		p.Iterations,
		p.Memory,
		p.Parallelism,
		p.KeyLength,
	)
	return hash, salt, nil
}

// Generates a random cryptographically secure byte slice.
//
// This function takes in an unsigned 32 bit integer to generate a byte slice 
// of the same length. 
// 
// It uses Go's built in `rand.Read` method to fill this byte slice
// accordingly. 
// 
// This is used in the salt generation step.
func generateSaltFromRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}
