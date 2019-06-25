package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"time"
)

func main() {

	listener, _ := net.Listen("tcp", ":8081")

	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		go handle(connection)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Fprintf(conn, "%s\tI heard you say: %s\n", time.Now().Format(time.RFC3339), text)
		if strings.ToUpper(text) == "SAIR" {
			break
		} else {
			conn.SetDeadline(time.Now().Add(10 * time.Second))
		}
	}
	defer conn.Close()
}
