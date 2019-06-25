package main

import (
	"fmt"
	"io"
	"log"
	"net"
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

		_, _ = io.WriteString(connection, "\nHello from tcp\n")
		_, _ = fmt.Fprintln(connection, "How are you?")
		_, _ = fmt.Fprintf(connection, "%v", "Well, I hope!\n")

		_ = connection.Close()
	}
}
