package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)

// Define our message object
type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

// Configure the upgrader
var upgrader = websocket.Upgrader{}

func main() {

	// Create a simple file server
	fs := http.FileServer(http.Dir("public"))
	http.Handle("/", fs)
}

func handleConnections(w http.ResponseWriter, r *http.Request) {

}

func handleMessage() {

}
