package gormkit

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid"`
	CreatedAt time.Time
	UpdatedAt time.Time

	Name string
	Age  int

	TriggeredAt time.Time
	CompletedAt *time.Time
	AbortedAt   *time.Time

	UniqueID *string `gorm:"uniqueIndex"`
	Test     bool
}
