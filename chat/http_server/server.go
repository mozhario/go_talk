package http_server

import (
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/mozhario/go_talk/config"
)

type HTTPServer struct {
	Mux *http.ServeMux
}

func (server *HTTPServer) Listen() {

	SetupRoutes(server.Mux)

	headers := handlers.AllowedHeaders([]string{"Content-Type", "Authorization"})
	origins := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"})

	handler := handlers.CORS(headers, origins, methods)(server.Mux)

	fmt.Println("HTTP server listening on " + config.ServerHost + ":" + config.ServerPort)
	http.ListenAndServe(fmt.Sprintf(":%s", config.ServerPort), handler)
}
