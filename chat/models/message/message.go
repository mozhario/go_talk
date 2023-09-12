package models

import (
	"time"
)

type Message struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Text     string
	Created  time.Time
	Updated  time.Time
}
