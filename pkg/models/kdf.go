package models

import (
	"crypto/rand"
	"fmt"
	"golang.org/x/crypto/argon2"
)

const (
	KDFAlgorithmV1Name   = "argon2id-v1"
	KDFAlgorithmV1Time   = 2
	KDFAlgorithmV1Memory = 64 * 1024
	KDFAlgorithmV1Thread = 4
	KDFAlgorithmV1KeyLen = 64
)

type KDF struct {
	Algorithm    string `gorm:"type:enum('argon2id-v1')"`
	PasswordHash [64]byte
	Salt         [64]byte
	TimeFactor   uint32
	MemoryFactor uint32
	ThreadFactor uint8
	KeyLen       uint32
}

func (k KDF) Validate(password string) bool {
	return false
}

func (k KDF) Derive(password string) [64]byte {
	panic("not implemented")
}

// NewKDF Generate a new KDF object without a password hash.
func NewKDF() KDF {
	var salt [64]byte

	count, err := rand.Read(salt[:])

	if count != 64 {
		panic("oh my GAWD")
	}

	if err != nil {
		panic(fmt.Sprintf("could not generate random salt: %s", err))
	}

	return KDF{
		Algorithm:    KDFAlgorithmV1Name,
		Salt:         salt,
		TimeFactor:   KDFAlgorithmV1Time,
		MemoryFactor: KDFAlgorithmV1Memory,
		ThreadFactor: KDFAlgorithmV1Thread,
		KeyLen:       KDFAlgorithmV1KeyLen,
	}
}

// GenKDF Generate a new KDF object using the given password, storing the password hash.
func GenKDF(password string) KDF {
	kdf := NewKDF()

	digest := argon2.IDKey([]byte(password), kdf.Salt[:], kdf.TimeFactor, kdf.MemoryFactor, kdf.ThreadFactor,
		kdf.KeyLen)

	copy(kdf.PasswordHash[:], digest)

	return kdf
}
