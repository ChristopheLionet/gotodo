package domain

import (
	"github.com/google/uuid"
)

type Task struct {
	Id   uuid.UUID
	Name string
}
