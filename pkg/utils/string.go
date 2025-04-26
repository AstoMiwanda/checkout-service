package utils

import "github.com/google/uuid"

func IsValidUUID(id uuid.UUID) bool {
	return id != uuid.Nil
}
