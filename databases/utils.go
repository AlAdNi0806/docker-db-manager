package databases

import (
	"fmt"
	"net"
	"os/exec"
)

func isDockerAvailable() bool {
	_, err := exec.LookPath("docker")
	return err == nil
}

func isPortInUse(port int) bool {
	// Attempt to listen on the port
	listener, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		// If there's an error (e.g., "address already in use"), the port is taken
		return true
	}
	// If we successfully listened, close and return false (port is free)
	listener.Close()
	return false
}
