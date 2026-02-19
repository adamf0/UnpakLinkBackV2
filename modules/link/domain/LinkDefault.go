package domain

import (
	"time"

	"github.com/google/uuid"
)

type LinkDefault struct {
	Id          uint
	UUID        uuid.UUID `gorm:"type:char(36);uniqueIndex"`
	ShortUrl    string
	LongUrl     string
	Creator     string
	StartAccess *time.Time
	EndAccess   *time.Time
	Password    *string
	Status      *string
}
