package websocket

import (
	"fmt"

	"github.com/mozhario/go_talk/chat/constants"
	models "github.com/mozhario/go_talk/chat/models/message"
	"github.com/mozhario/go_talk/chat/services"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan models.Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan models.Message),
	}
}

func (pool *Pool) Start() {
	for {
		select {
		case client := <-pool.Register:
			pool.register(client)
			break
		case client := <-pool.Unregister:
			pool.unregister(client)
			break
		case message := <-pool.Broadcast:
			pool.broadcastMessage(message)
		}
	}
}

func (pool *Pool) register(client *Client) {
	pool.Clients[client] = true
	fmt.Println("Size of Connection Pool: ", len(pool.Clients))
	message := models.Message{
		Type:     constants.MessageTypeSystemMessage,
		Username: constants.MessageTypeSystemMessage,
		Text:     constants.CientJoined,
	}

	services.SaveMessage(message)

	for client, _ := range pool.Clients {
		client.Conn.WriteJSON(message)
	}
}

func (pool *Pool) unregister(client *Client) {
	delete(pool.Clients, client)
	fmt.Println("Size of Connection Pool: ", len(pool.Clients))

	message := models.Message{
		Type:     constants.MessageTypeSystemMessage,
		Username: constants.MessageTypeSystemMessage,
		Text:     constants.ClientDisconnected,
	}

	services.SaveMessage(message)

	for client, _ := range pool.Clients {
		client.Conn.WriteJSON(message)
	}
}

func (pool *Pool) broadcastMessage(message models.Message) {
	fmt.Println("Broadcasting message: ", message)

	services.SaveMessage(message)

	for client, _ := range pool.Clients {
		if err := client.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
			return
		}
	}
}
