package clientserve

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"

	"netcat/utils"
)


func HandleConnection(conn net.Conn) {
	if len(utils.Clients) >= utils.MaxClients {
		fmt.Fprint(conn, "There are too many connections, try to log in later.")
		conn.Close()
		return
	}

	log.Println("Starting to handle connection. Current Clients:", len(utils.Clients))

	var name string
	var err error
	// Read client's name
	reader := bufio.NewReader(conn)

	// Send welcome message
	utils.WelcomeMessage(conn)

	for {
		fmt.Fprint(conn, "\n[ENTER YOUR NAME]: ")
		name, err = reader.ReadString('\n')
		if err != nil {
			log.Printf("Error reading client name: %v", err)
			utils.ClientExit(conn)
			return
		}
		name = strings.TrimSpace(name)
		name, err = utils.CheckText(name, "Checkname")
		if err != nil {
			fmt.Fprint(conn, err)
		} else if name == "" {
			fmt.Fprint(conn, "Can't send empty name\n")
		} else if len(name) >= 10 {
			fmt.Fprint(conn, "Max characters: 9\n")
		} else if utils.AlreadyExist(name) {
			fmt.Fprint(conn, "Name have been taken , try another name\n")
		} else {
			break
		}
	}

	defer utils.ClientExit(conn)
	utils.AddClient(conn, name)

	// Send previous messages to the new client
	utils.SendPreviousMessages(conn)

	// system.Broadcast that a new client has joined
	utils.Broadcast(fmt.Sprintf("%s has joined the chat.\n", name), conn)

	// Handle messages from the client
	for {
		// Prompt the client for a message
		fmt.Fprintf(conn, "[%s][%s]: ", time.Now().Format("2006-01-02 15:04:05"), name)

		message, err := reader.ReadString('\n')
		if err != nil {
			if !strings.Contains(err.Error(), "use of closed network connection") {
				log.Printf("Error reading from client: %v", err)
			}
			return
		}

		message = strings.TrimSpace(message)
		message, err = utils.CheckText(message, "Checkmessage")
		if err != nil {
			fmt.Fprint(conn, err)
			fmt.Fprintln(conn)
			continue
		} else if message == "" {
			fmt.Fprint(conn, "Can't send empty messages\n")
			continue
		} else if strings.HasPrefix(message, "--name") {
			newName := strings.TrimSpace(message[len("--name"):])
			newName, err = utils.CheckText(newName, "Checkname")
			if err != nil || newName == "" || len(newName) >= 10 {
				fmt.Fprint(conn, "Invalid new name. Please use a valid name (1-9 characters).\n")
				continue
			}

			oldName := utils.Clients[conn]
			utils.Mutex.Lock()
			utils.Clients[conn] = newName
			name = newName
			utils.Mutex.Unlock()

			utils.Broadcast(fmt.Sprintf("%s has changed their name to %s.\n", oldName, newName), conn)
			log.Printf("Client : '%s' has changed their name to '%s'.\n", oldName, newName)
			fmt.Fprint(conn, "Your name has been changed successfully.\n")
			continue
		}

		// Format the message with timestamp and sender's name
		timestamp := time.Now().Format("[2006-01-02 15:04:05]")
		formattedMessage := fmt.Sprintf("%s[%s]: %s\n", timestamp, name, message)

		// utils.Broadcast the message to all other Clients
		utils.Broadcast(formattedMessage, conn)
	}
}
