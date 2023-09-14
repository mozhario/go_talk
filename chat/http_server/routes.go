package http_server

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/mozhario/go_talk/chat/services"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", root)
	mux.HandleFunc("/messages", messagesList)
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("root")
	b, err := ioutil.ReadFile("static/root.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(b))
}

func messagesList(w http.ResponseWriter, r *http.Request) {
	messages, err := services.RetrieveMessagesFromDatabase()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}
