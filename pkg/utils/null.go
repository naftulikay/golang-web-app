package utils

func NullBytes(length uint) []byte {
	return make([]byte, length)
}

func NullBytes32() [32]byte {
	return [32]byte{}
}

func NullBytes64() [64]byte {
	return [64]byte{}
}
