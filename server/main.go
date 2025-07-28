package main

import (
	"fmt"
	"log"
	"net"
	"os"

	clientserve "netcat/client"
	"netcat/utils"
)

func main() {
	logFile, err := utils.SetupLogging()
	if err != nil {
		fmt.Printf("Failed to set up logging: %v\n", err)
		return
	}
	defer logFile.Close()

	// default port
	port := "8989"

	// check if a port is provided as a command-line argument
	if len(os.Args) > 1 {
		if len(os.Args) != 2 {
			if os.Args[2] == "--clear" {
				utils.Truncate(logFile)
			} else {
				fmt.Println("[USAGE]: ./TCPChat $port")
				return
			}
		}
		// check clear logs flag
		if os.Args[1] == "--clear" {
			utils.Truncate(logFile)
		} else {
			port = os.Args[1]
		}
	}

	// Start listening on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("Error listening: %v", err)
	}
	defer listener.Close()

	log.Printf("Listening on port: %s\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		log.Printf("New connection attempt. Current Clients: %d\n", len(utils.Clients))

		go clientserve.HandleConnection(conn)
	}
}
