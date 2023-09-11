package server

import (
	"fmt"
	// "log"

	"net/http"
	"os"

	"github.com/gorilla/websocket"
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
	// Infinite loop to continuously read messages from the client
	for {
		// Read message from the client
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err.Error())
			break
		}

		// Print the received message
		fmt.Printf("Received message: %s\n", message)

		// Send a response back to the client
		err = conn.WriteMessage(websocket.TextMessage, []byte("Message received."))
		if err != nil {
			fmt.Println("Error writing message:", err.Error())
			break
		}
	}

	// Close the connection when done
	conn.Close()
}

// func (server WebSocketServer) HandleRequest(conn net.Conn) {
// 	// Make a buffer to hold incoming data.
// 	buf := make([]byte, 1024)
// 	// Read the incoming connection into the buffer.
// 	_, err := conn.Read(buf)
// 	if err != nil {
// 		fmt.Println("Error reading:", err.Error())
// 	}
// 	// Send a response back to person contacting us.
// 	conn.Write([]byte("Message received."))
// 	// Close the connection when you're done with it.
// 	conn.Close()
// }
