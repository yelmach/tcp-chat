# TCP Chat - NetCat

A Go implementation of a TCP-based group chat server inspired by NetCat, supporting multiple concurrent clients with real-time messaging, connection management, and comprehensive logging.

## Overview

This project recreates the NetCat functionality in a Server-Client architecture, allowing multiple clients to connect to a TCP server and participate in a group chat. The server manages connections, broadcasts messages, and maintains chat history for new joining clients.

## Features

### Core Functionality
- **TCP Server-Client Architecture**: One server handling multiple concurrent clients
- **Real-time Group Chat**: Instant message broadcasting to all connected clients
- **Connection Management**: Maximum 10 concurrent connections with graceful handling
- **Message History**: New clients receive all previous chat messages
- **User Authentication**: Name requirement and validation for all clients
- **Timestamped Messages**: All messages include timestamp and sender identification

### Advanced Features
- **Name Change Support**: Clients can change their names using `--name` command
- **Join/Leave Notifications**: Server announces when clients join or leave
- **Input Validation**: Comprehensive text validation and error handling
- **Logging System**: Server logs all activities to file with optional log clearing
- **Connection Limits**: Prevents server overload with maximum connection enforcement

### Technical Implementation
- **Go Routines**: Concurrent handling of multiple client connections
- **Mutexes**: Thread-safe access to shared resources
- **Channels**: Efficient communication between goroutines
- **Error Handling**: Robust error management for network operations

## Usage

### Starting the Server

Default port (8989):
```bash
go run server/main.go
```

Custom port:
```bash
go run server/main.go 2525
```

Clear logs:
```bash
go run server/main.go --clear
go run server/main.go 2525 --clear
```

### Connecting Clients

Using netcat:
```bash
nc localhost 8989
```

Using telnet:
```bash
telnet localhost 8989
```

## Client Interaction Flow

1. **Connection**: Client connects and receives welcome message with Linux logo
2. **Name Entry**: Client enters a unique name (1-9 characters, no spaces)
3. **Chat History**: Client receives all previous messages
4. **Real-time Chat**: Client can send/receive messages in real-time
5. **Name Change**: Client can change name using `--name newname`
6. **Disconnection**: Clean exit with notification to other clients

### Example Session

```
Welcome to TCP-Chat!
         _nnnn_
        dGGGGMMb
       @p~qp~~qMb
       M|@||@) M|
       @,----.JM|
      JS^\__/  qKL
     dZP        qKRb
    dZP          qKKb
   fZP            SMMb
   HZM            MMMM
   FqM            MMMM
 __| ".        |\dS"qML
 |    `.       | `' \Zq
_)      \.___.,|     .'
\____   )MMMMMP|   .'
     `-'       `--'

[ENTER YOUR NAME]: Alice
[2024-01-20 15:48:41][Alice]: Hello everyone!

Bob has joined the chat.
[2024-01-20 15:49:15][Bob]: Hi Alice!
[2024-01-20 15:49:20][Alice]: Welcome Bob!
```

## Message Format

All messages follow the format:
```
[YYYY-MM-DD HH:MM:SS][username]: message
```

System notifications:
```
username has joined the chat.
username has left the chat.
username has changed their name to newname.
```

## Project Structure

```
tcp-chat/
├── server/
│   ├── main.go              # Server entry point and connection handling
│   └── server.log           # Server activity logs
├── client/
│   └── clients.go           # Client connection and message handling
├── utils/
│   ├── globalvariables.go   # Shared variables and data structures
│   ├── Validation.go        # Input validation functions
│   ├── change.go            # Client management (add/remove/exit)
│   ├── send.go              # Message broadcasting and history
│   ├── welcome.go           # Welcome message and Linux logo
│   └── logging.go           # Logging setup and management
├── go.mod                   # Go module definition
└── README.md
```

## Validation Rules

### Name Validation
- **Length**: 1-9 characters maximum
- **Characters**: No spaces or control characters (ASCII < 32)
- **Uniqueness**: Names must be unique across all connected clients
- **Required**: Empty names are not allowed

### Message Validation
- **Content**: No control characters (ASCII < 32) except spaces
- **Empty Messages**: Cannot send empty messages
- **Special Commands**: `--name newname` for name changes

## Error Handling

The server handles various error conditions:
- **Connection Limit**: "There are too many connections, try to log in later."
- **Invalid Characters**: "Invalid characters in your input"
- **Empty Input**: "Can't send empty name/messages"
- **Name Conflicts**: "Name have been taken, try another name"
- **Network Errors**: Graceful disconnection and cleanup

## Logging

Server maintains comprehensive logs in `server.log`:
- Client connections and disconnections
- Name changes and validations
- Error conditions and network issues
- Server start/stop events

Log management:
```bash
# Clear logs on startup
go run server/main.go --clear

# View logs
tail -f server.log
```

## Technical Details

### Concurrency
- **Go Routines**: Each client connection runs in a separate goroutine
- **Mutexes**: Protect shared data structures (client map, message history)
- **Thread Safety**: All operations on shared resources are mutex-protected

### Network Programming
- **TCP Sockets**: Reliable connection-oriented communication
- **Connection Pooling**: Efficient management of multiple connections
- **Graceful Shutdown**: Clean disconnection handling

### Data Structures
```go
var Clients = make(map[net.Conn]string)  // Connection to name mapping
var Messages []string                     // Chat history storage
var Mutex sync.Mutex                     // Thread synchronization
```
