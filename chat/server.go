package server

import (
	"fmt"
	"log"
	"net"

	"github.com/mozhario/go_talk/constants"
	"github.com/mozhario/go_talk/settings"
)

type Server struct {
	Port string
}

func (server Server) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprint(":%s", server.Port))
	if err != nil {
		log.Fatal(constants.ErrorListenerFailedToStart)
	}

	conn, _ := listener.Accept()
	conn.Write([]bytes(constants.ChatStarted))

	for {
		buff := make([]byte, settings.ServerBufferSize)
		msg, err := conn.Read(buff[0:])
		if err != nil {
			log.Fatal(constants.ErrorReadingMessage)
		}
		conn.Write(msg)
	}
}
