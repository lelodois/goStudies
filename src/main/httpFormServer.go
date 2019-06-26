package main

import (
	"fmt"
	"net/http"
)

type handleValue int

func (m handleValue) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseForm()
	fmt.Fprintf(writer, "Form with values: %s", request.Form.Encode())
}

func main() {
	var varHandle handleValue
	http.ListenAndServe(":8081", varHandle)
}
