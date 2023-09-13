package websocket

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
	"github.com/mozhario/go_talk/chat/constants"
	models "github.com/mozhario/go_talk/chat/models/message"
	"github.com/mozhario/go_talk/db"
)

type Client struct {
	ID   string
	Conn *websocket.Conn
	Pool *Pool
}

func (client *Client) Read() {
	defer func() {
		client.Pool.Unregister <- client
		client.Conn.Close()
	}()

	for {
		_, msg, err := client.Conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}

		var parsedMessage Message
		if err := client.parseMessage(msg, &parsedMessage); err != nil {
			fmt.Println("Error parsing JSON:", err.Error())
			break
		}

		client.SaveMessage(parsedMessage)

		client.Pool.Broadcast <- parsedMessage
		fmt.Printf("Message Received: %+v\n", parsedMessage)
	}
}

func (client Client) SaveMessage(messageData Message) {
	message := models.Message{
		Username: messageData.Username,
		Text:     messageData.Text,
	}
	db.DB.Create(&message)
}

func (client Client) parseMessage(messageData []byte, message *Message) error {
	err := json.Unmarshal(messageData, &message)
	message.Type = constants.MessageTypeUserMessage
	return err
}
