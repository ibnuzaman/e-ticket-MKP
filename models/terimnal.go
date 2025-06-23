package models

import (
	"time"

	"github.com/google/uuid"
)

type Terminal struct {
	TerminalID uuid.UUID `gorm:"type:uuid;primaryKey"`
	Name       string    `gorm:"type:varchar(100);not null"`
	Location   string    `gorm:"type:varchar(255);not null"`
	Status     string    `gorm:"type:varchar(20);not null"`
	CreatedAt  time.Time `gorm:"not null;default:current_timestamp"`
	UpdatedAt  time.Time `gorm:"not null;default:current_timestamp"`
}
