package routes

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	log.Println("root")
	b, err := ioutil.ReadFile("static/root.txt")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(w, string(b))
}
