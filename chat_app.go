package main

import (
	"net/http"
)

// Define our message object
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

}

func handleMessage() {

}
