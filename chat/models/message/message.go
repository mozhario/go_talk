package models

import (
	"time"
)

type Message struct {
	ID       uint      `gorm:"primaryKey" json:"id"`
	Type     string    `json:"type"`
	Username string    `json:"username"`
	Text     string    `json:"text"`
	Created  time.Time `json:"created"`
	Updated  time.Time `json:"updated"`
	SentTime time.Time `json:"sent_time"`
}
