package utils

import (
	"net"
	"sync"
)

var MaxClients = 10

var Clients = make(map[net.Conn]string)
var (
	Mutex    sync.Mutex
	Messages []string
)
