package dto

import "github.com/google/uuid"

type User struct {
	ID   uuid.UUID
	Name string
	Age  int
}

