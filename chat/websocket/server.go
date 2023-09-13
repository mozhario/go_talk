package websocket

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

type WebSocketServer struct {
	Host string
	Port string
	Pool *Pool
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func (server WebSocketServer) Listen() {
	go server.Pool.Start()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading:", err.Error())
			return
		}
		server.HandleRequest(conn)
	})

	fmt.Println("WebSocket server listening on " + server.Host + ":" + server.Port)
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", server.Host, server.Port), nil)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
}

func (server WebSocketServer) HandleRequest(conn *websocket.Conn) {
	defer conn.Close()

	client := &Client{
		Conn: conn,
		Pool: server.Pool,
	}

	server.Pool.Register <- client
	client.Read()
}
