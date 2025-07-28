package utils

import "fmt"

func CheckText(text string, flag string) (s string, err error) {
	for _, Char := range text {
		if flag == "Checkname" {
			if Char < 32 || Char == ' ' {
				return "", fmt.Errorf("invalid characters in your input")
			}
		} else if flag == "Checkmessage" {
			if Char < 32 {
				return "", fmt.Errorf("invalid characters in your input")
			}
		}
	}
	return text, nil
}

func AlreadyExist(name string) bool{
	Mutex.Lock()
	defer Mutex.Unlock()
	for _, clientName := range Clients {
		if clientName == name {
			return true
		}
	}
	return false
}
