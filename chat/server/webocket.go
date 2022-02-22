package server

import (
	"fmt"
	"log"
	"net"

	"github.com/mozhario/go_talk/config/settings"
	"github.com/mozhario/go_talk/constants/messages"
)

type WebSocketServer struct {
	Port string
}

func (server WebSocketServer) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprint(":%s", server.Port))
	if err != nil {
		log.Fatal(messages.ErrorListenerFailedToStart)
	}

	conn, _ := listener.Accept()
	conn.Write([]byte(messages.ChatStarted))

	for {
		buff := make([]byte, settings.ServerBufferSize)
		msg, err := conn.Read(buff[0:])
		if err != nil {
			log.Fatal(messages.ErrorReadingMessage)
		}
		conn.Write(msg)
	}
}
