package main

import (
	"log"
	"net/http"

	"github.com/papudatta/webrocket/internal/handlers"
)

func main() {
	mux := routes()

	log.Println("Starting channel listener ..")
	go handlers.ListenToWsChannel()

	log.Println("Starting web server on port 9090")

	_ = http.ListenAndServe(":9090", mux)
}
