package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
	"github.com/mozhario/go_talk/chat/db"
	models "github.com/mozhario/go_talk/chat/models/message"
)

type WebSocketServer struct {
	Host string
	Port string
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (server WebSocketServer) Listen() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading:", err.Error())
			return
		}
		server.HandleRequest(conn)
	})

	fmt.Println("Listening on " + server.Host + ":" + server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.Host, server.Port), nil)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
}

func (server WebSocketServer) HandleRequest(conn *websocket.Conn) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}

		fmt.Printf("Received message: %s\n", message)

		server.SaveMessage(message)

		err = conn.WriteMessage(websocket.TextMessage, []byte(message))
		if err != nil {
			fmt.Println("Error writing message:", err.Error())
			break
		}
	}
}

func (server WebSocketServer) SaveMessage(messageData []byte) error {
	var message models.Message

	err := json.Unmarshal(messageData, &message)
	if err != nil {
		fmt.Println("Error parsing JSON:", err.Error())
		return err
	}
	db.DB.Create(&message)
	return err
}
