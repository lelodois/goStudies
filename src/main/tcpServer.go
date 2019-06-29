package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"reflect"
	"strings"
	"time"
)

var repository = make(map[string]string)

func main() {

	listener, _ := net.Listen("tcp", ":8081")

	var instructions = []instruction{del{}, set{}, getByValue{}, getAll{}, get{}}

	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		go handle(connection, instructions)

	}
}

type instruction interface {
	do(conn net.Conn, parameters []string)
}

type set struct {
}

type get struct {
}

type getByValue struct {
}

type getAll struct {
}

type del struct {
}

func (d get) do(conn net.Conn, parameters []string) {
	if len(parameters) != 2 {
		fmt.Fprintf(conn, "Expected command and key, for example: '%s name'\n", commandName(d))
	} else {
		value := repository[parameters[1]]
		fmt.Fprintf(conn, "%s\n", value)
	}
}

func (d getByValue) do(conn net.Conn, parameters []string) {
	if len(parameters) != 2 {
		fmt.Fprintf(conn, "Expected command and value, for example: '%s Léo'\n", commandName(d))
	} else {
		for key := range repository {
			if repository[key] == parameters[1] {
				fmt.Fprintf(conn, "%s\n", key)
			}
		}
	}
}

func (d getAll) do(conn net.Conn, parameters []string) {
	if len(parameters) != 1 {
		fmt.Fprintf(conn, "Expected command, for example: '%s'\n", commandName(d))
	} else {
		fmt.Fprintf(conn, "%s\n", repository)
	}
}

func (d set) do(conn net.Conn, parameters []string) {
	if len(parameters) != 3 {
		fmt.Fprintf(conn, "Expected command, key and value, for example: '%s name Léo'\n", commandName(d))
	} else {
		repository[parameters[1]] = parameters[2]
	}
}

func (d del) do(conn net.Conn, parameters []string) {
	if len(parameters) != 2 {
		fmt.Fprintf(conn, "Expected command and key, for example: '%s name'\n", commandName(d))
	} else {
		delete(repository, parameters[1])
	}
}

func commandName(instance instruction) string {
	return strings.ToLower(reflect.TypeOf(instance).Name())
}

func handle(conn net.Conn, instructions []instruction) {
	scanner := bufio.NewScanner(conn)
	io.WriteString(conn,
		"Use commands:\n"+
			"\t'set key value'\n"+
			"\t'get key'\n"+
			"\t'getbyvalue value'\n"+
			"\t'getall'\n"+
			"\t'del key value'\n"+
			"\t'exit'\n")

	for scanner.Scan() {
		parameters := strings.Fields(scanner.Text())
		command := strings.ToLower(parameters[0])
		if len(parameters) == 0 || command == "exit" {
			fmt.Fprintf(conn, "Good bye\n")
			conn.Close()
		} else {
			conn.SetDeadline(time.Now().Add(15 * time.Second))
			fmt.Fprintf(conn, "%s\treceived: %s\n", time.Now().Format(time.RFC3339), parameters)

			for _, item := range instructions {
				if commandName(item) == command {
					item.do(conn, parameters)
				}
			}
		}
	}
	defer conn.Close()
}
