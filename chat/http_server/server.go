package http_server

import (
	"fmt"
	"net/http"

	"github.com/mozhario/go_talk/config"
)

type HTTPServer struct {
	Mux *http.ServeMux
}

func (server *HTTPServer) Listen() {

	SetupRoutes(server.Mux)
	fmt.Println("HTTP server listening on " + config.ServerHost + ":" + config.ServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), server.Mux)
}
