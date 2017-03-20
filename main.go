package main

import (
	"./server"
	"os/exec"
	"fmt"
)

// TODO: Test
func main() {
	s := server.New()
	go s.Execute()

	// TODO: Move to server package
	cmd := "prometheus -config.file=/etc/prometheus/prometheus.yml -storage.local.path=/prometheus -web.console.libraries=/usr/share/prometheus/console_libraries -web.console.templates=/usr/share/prometheus/consoles"
	// TODO: Output to stdout/stderr
	_, err := exec.Command("/bin/sh", "-c", cmd).CombinedOutput()
	if err != nil {
		fmt.Printf("Could not execute the command:\n%s\n\n%s\n", cmd, err.Error())
	}
}
