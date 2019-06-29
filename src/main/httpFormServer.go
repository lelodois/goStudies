package main

import (
	"fmt"
	"net/http"
	"os"
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

func faviIconServerHttp(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "text/html; charset=utf-8")
	file, err := os.Open("/favicon.ico")
	if err != nil {
		http.Error(response, "File favicon not found", 404)
		return
	}
	defer file.Close()

	info, _ := file.Stat()

	http.ServeContent(response, request, file.Name(), info.ModTime(), file)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/user", userServerHttp)
	mux.HandleFunc("/admin", adminServerHttp)
	mux.HandleFunc("/icon", faviIconServerHttp)

	http.ListenAndServe(":8081", mux)
}
