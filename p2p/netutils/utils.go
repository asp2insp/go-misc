package netutils

import (
	"fmt"
	"net"
)

// GetLocalAddress returns the local IP address
// concatenated with the given port
func GetLocalAddress(port int) string {
	address := getLocalIP()
	return fmt.Sprintf("http://%s:%d", address, port)
}

// GetLocalIP returns the non loopback local IP of the host
func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
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
