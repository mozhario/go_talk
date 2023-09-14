package services

import (
	"time"

	models "github.com/mozhario/go_talk/chat/models/message"
	"github.com/mozhario/go_talk/db"
)

func RetrieveMessagesFromDatabase() ([]models.Message, error) {
	var messages []models.Message
	result := db.DB.Find(&messages)
	if result.Error != nil {
		return nil, result.Error
	}
	return messages, nil
}

func SaveMessage(message models.Message) {
	message.Created = time.Now()
	message.Updated = time.Now()
	db.DB.Create(&message)
}
