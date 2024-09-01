package utils

import "github.com/google/uuid"

func GenerateUUID() string {
	id, _ := uuid.NewV7()
	return id.String()
}
