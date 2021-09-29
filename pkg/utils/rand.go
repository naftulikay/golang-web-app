package utils

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
)

func InfallibleSecureRandBase64(length uint) string {
	data := make([]byte, length)

	count, err := rand.Read(data)

	if err != nil {
		panic(fmt.Sprintf("unable to read from random number generator: %s", err))
	}

	if uint(count) != length {
		panic(fmt.Sprintf("tried to read %d bytes, but only read %d bytes from CSPRNG", length, count))
	}

	return base64.URLEncoding.EncodeToString(data)
}

func InfallibleSecureRandBytes32() [32]byte {
	var data [32]byte

	count, err := rand.Read(data[:])

	if err != nil {
		panic(fmt.Sprintf("unable to read from random number generator: %s", err))
	}

	if count != 32 {
		panic(fmt.Sprintf("tried to read %d bytes, but only read %d bytes from CSPRNG", 32, count))
	}

	return data
}
