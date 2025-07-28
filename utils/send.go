package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

func Broadcast(message string, sender net.Conn) {
	Mutex.Lock()
	defer Mutex.Unlock()

	// Store the message
	Messages = append(Messages, message)

	for client := range Clients {
		if client != sender {
			_, err := client.Write([]byte("\n" + message))
			fmt.Fprintf(client, "[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), Clients[client])
			if err != nil {
				log.Printf("Error broadcasting to client: %v", err)
				ClientExit(client)
				RemoveClient(client)
			}
		}
	}
}

func SendPreviousMessages(conn net.Conn) {
	Mutex.Lock()
	defer Mutex.Unlock()

	for _, message := range Messages {
		_, err := conn.Write([]byte(message))
		if err != nil {
			log.Printf("Error sending previous message to client: %v", err)
			return
		}
	}
}
