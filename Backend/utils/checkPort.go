package utils

import (
	"net"
)

// Check if a port is available
func CheckPort(port string) bool {
	ln, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return false
	}
	ln.Close()
	return true
}
