package websocket

import (
	"fmt"

	"github.com/mozhario/go_talk/chat/constants"
)

type Pool struct {
	Register   chan *Client
	Unregister chan *Client
	Clients    map[*Client]bool
	Broadcast  chan Message
}

func NewPool() *Pool {
	return &Pool{
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
		Clients:    make(map[*Client]bool),
		Broadcast:  make(chan Message),
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
	for client, _ := range pool.Clients {
		// fmt.Println(client)
		client.Conn.WriteJSON(
			Message{
				Type:     constants.MessageTypeSystemMessage,
				Username: constants.MessageTypeSystemMessage,
				Text:     constants.CientJoined,
			},
		)
	}
}

func (pool *Pool) unregister(client *Client) {
	delete(pool.Clients, client)
	fmt.Println("Size of Connection Pool: ", len(pool.Clients))
	for client, _ := range pool.Clients {
		client.Conn.WriteJSON(
			Message{
				Type:     constants.MessageTypeSystemMessage,
				Username: constants.MessageTypeSystemMessage,
				Text:     "User Disconnected...",
			},
		)
	}
}

func (pool *Pool) broadcastMessage(message Message) {
	fmt.Println("Broadcasting message: ", message)
	for client, _ := range pool.Clients {
		if err := client.Conn.WriteJSON(message); err != nil {
			fmt.Println(err)
			return
		}
	}
}
