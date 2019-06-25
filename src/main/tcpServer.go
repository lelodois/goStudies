package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"
)

var repository = make(map[string]string)

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
	instructions :=
		"Use commands:\n" +
			"\t'SET key value'\n" +
			"\t'GET_BY_KEY key'\n" +
			"\t'GET_BY_VALUE value'\n" +
			"\t'GET_ALL'\n" +
			"\t'DEL key value'\n" +
			"\t'EXIT'\n"
	io.WriteString(conn, instructions)

	for scanner.Scan() {
		conn.SetDeadline(time.Now().Add(10 * time.Second))

		text := scanner.Text()
		fmt.Fprintf(conn, "%s\treceived: %s\n", time.Now().Format(time.RFC3339), text)
		parameters := strings.Fields(text)

		switch strings.ToUpper(parameters[0]) {

		case "GET_BY_KEY":
			if len(parameters) != 2 {
				fmt.Fprintf(conn, "For GET_BY_KEY - Expected command and key, for example: 'GET_BY_KEY name'\n")
				continue
			}
			value := repository[parameters[1]]
			fmt.Fprintf(conn, "%s\n", value)

		case "GET_BY_VALUE":
			if len(parameters) != 2 {
				fmt.Fprintf(conn, "For GET_BY_VALUE - Expected command and value, for example: 'GET_BY_VALUE Léo'\n")
				continue
			}
			for key := range repository {
				if repository[key] == parameters[1] {
					fmt.Fprintf(conn, "%s\n", key)
				}
			}

		case "GET_ALL":
			if len(parameters) != 1 {
				fmt.Fprintf(conn, "For GET_ALL - Expected command, for example: 'GET_ALL'\n")
				continue
			}
			if len(repository) != 0 {
				fmt.Fprintf(conn, "%s\n", repository)
			}

		case "SET":
			if len(parameters) != 3 {
				fmt.Fprintf(conn, "For SET - Expected command, key and value, for example: 'SET name Léo'\n")
				continue
			}
			repository[parameters[1]] = parameters[2]

		case "DEL":
			if len(parameters) != 2 {
				fmt.Fprintf(conn, "for DEL - Expected command and key, for example: 'DEL name'\n")
				continue
			}
			delete(repository, parameters[1])
		case "EXIT":
			fmt.Fprintf(conn, "Good bye\n")
			conn.Close()
			break
		default:
			fmt.Fprintf(conn, "Invalid command \n%s", instructions)
		}
	}

	defer conn.Close()
}
