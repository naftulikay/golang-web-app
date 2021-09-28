package utils

// Nop Do exactly nothing, useful to not break compilation on variables which do need to exist in spite of not being
// used.
func Nop(args ...interface{}) {}
