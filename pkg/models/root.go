package models

// Models List all known database models.
func Models() []interface{} {
	return []interface{}{
		User{},
	}
}
