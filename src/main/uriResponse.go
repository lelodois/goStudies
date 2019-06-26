package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

func main() {
	listener, _ := net.Listen("tcp", ":8080")

	defer listener.Close()

	for {
		connection, _ := listener.Accept()
		go request(connection)

	}
}

func request(connection net.Conn) {
	defer connection.Close()

	scanner := bufio.NewScanner(connection)
	index := 0
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Println(text)
		if index == 0 {
			mux(connection, text)
		}
		if text == "" {
			break
		}
		index++
	}
}

func mux(connection net.Conn, text string) {
	method := strings.Fields(text)[0]
	uri := strings.Fields(text)[1]

	if method == "GET" && uri == "/" {
		navigate(connection, "Home Page")
	}

	if method == "GET" && uri == "/admin" {
		navigate(connection, "Admin Page")
	}

	if method == "GET" && uri == "/login" {
		navigate(connection, "Login Page")
	}
}

func navigate(conn net.Conn, text string) {
	html := `
		<!DOCTYPE html>
		<html lang="en">
			<head><meta charset="UTF-8"><title>Hello Uris!</title></head>
			<body>
				<div style="margin-bottom:50px;"><h3>Você está em: ` + text + `</h3></div>
				<ul>
					<li><a href="/">Home</a><br></li>
					<li><a href="/admin">Admin</a><br></li>
					<li><a href="/login">Login</a><br></li>
				</ul>
			</body>
		</html>
	`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(html))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, html)
}
