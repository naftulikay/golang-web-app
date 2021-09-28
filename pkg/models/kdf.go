package models

import (
	"crypto/rand"
	"fmt"
)

type KDF struct {
	Algorithm    string
	PasswordHash [64]byte
	Salt         [64]byte
	Compute      uint32
	Memory       uint32
	Concurrency  uint8
	KeyLen       uint32
}

func (k KDF) Validate(password string) bool {
	return false
}

func (k KDF) Derive(password string) [64]byte {
	panic("not implemented")
}

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
		Algorithm:   "argon2id-v1",
		Salt:        salt,
		Compute:     2,
		Memory:      64 * 1024,
		Concurrency: 4,
		KeyLen:      64,
	}
}
