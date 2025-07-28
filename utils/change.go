package utils

import (
	"fmt"
	"log"
	"net"
)

func AddClient(conn net.Conn, name string) {
	Mutex.Lock()
	defer Mutex.Unlock()

	Clients[conn] = name
	log.Printf("Client connected : %s. Total Clients: %d\n", name, len(Clients))
}

func RemoveClient(client net.Conn) {
	Mutex.Lock()
	defer Mutex.Unlock()

	name := Clients[client]
	delete(Clients, client)
	log.Printf("Client disconnected : %s. Total Clients: %d\n", name, len(Clients))
}

func ClientExit(conn net.Conn) {
	Broadcast(fmt.Sprintf("%s has left the chat.\n", Clients[conn]), conn)
	RemoveClient(conn)
	conn.Close()
}
