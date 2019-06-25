package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	listener, serverError := net.Listen("tcp", ":8081")

	if serverError != nil {
		log.Panic(serverError)
	}

	defer listener.Close()

	for {
		connection, connectionError := listener.Accept()

		if connectionError != nil {
			log.Println(connectionError)
		}
		go handle(connection)

	}
}

func handle(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		fmt.Fprintf(conn, "I heard you say: %s\n", text)
		if strings.ToUpper(text) == "SAIR" {
			break
		}
	}
	defer conn.Close()
}
