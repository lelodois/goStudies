package main

import (
	"fmt"
	"net/http"
)

func adminServerHttp(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.Header().Set("User-Type", "admin")
	fmt.Fprintf(response, "Save Admin user with values: %s", request.Form.Encode())
}

func userServerHttp(response http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	response.Header().Set("User-Type", "user")
	fmt.Fprintf(response, "Save simple user with values: %s", request.Form.Encode())
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userServerHttp)
	mux.HandleFunc("/admin", adminServerHttp)

	http.ListenAndServe(":8081", mux)
}
