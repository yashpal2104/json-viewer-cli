// /*
// Copyright © 2025 NAME HERE <EMAIL ADDRESS>
// */
package main

import (
	"log"

	"github.com/gorilla/websocket"
	"github.com/yashpal2104/json-viewer-cli/cmd"
	// "github.com/yashpal2104/json-viewer-cli/data"
)

func main() {
	cmd.Execute()
	dummyUse()
}

func dummyUse() {
	_, _, err := websocket.DefaultDialer.Dial("ws://localhost:8080/ws", nil)
	if err != nil {
		log.Println("Dummy dial error:", err)
	}
}
