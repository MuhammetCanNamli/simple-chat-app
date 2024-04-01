package main

import (
	"fmt"
)

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	fmt.Println("ChatApp with WebSocket")
}
