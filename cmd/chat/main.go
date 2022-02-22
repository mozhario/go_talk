package main

import (
	"github.com/mozhario/go_talk/chat/server/websocket"
	"github.com/mozhario/go_talk/config/settings"
)

func main() {
	server := websocket.WebSocketServer{settings.ServerPort}
	server.Listen()
}
