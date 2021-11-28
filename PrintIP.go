package main

import (
	"fmt"
	"net"
)

// on Mac you can check by running from terminal : ifconfig | grep "inet " | grep -v 127.0.0.1 | cut -d\  -f2
// GetLocalIP returns the non loopback local IP of the host
func GetLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "Error fetching network interfaces..."
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func main() {
	fmt.Printf(" My IP is %v", GetLocalIP())
}
