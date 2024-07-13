package question

import (
	"time"

	"gorm.io/gorm"
)

type Question struct {
	ID            uint `gorm:"primarykey"`
	Type          string
	Text          string
	CorrectAnswer string
	Options       string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}
