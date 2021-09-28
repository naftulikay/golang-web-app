package service

type JWTServiceImpl struct {
	key [32]byte
}

func (J JWTServiceImpl) Validate(token string) bool {
	return false
}

func NewJWTService(key [32]byte) JWTServiceImpl {
	return JWTServiceImpl{key: key}
}
