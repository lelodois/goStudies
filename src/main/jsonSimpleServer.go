package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Message struct {
	Id       int64  `json:"Id"`
	Name     string `json:"Name"`
	LastName string `json:"lastname"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var m Message
		_ = json.NewDecoder(r.Body).Decode(&m)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json, _ := json.Marshal(m)
		fmt.Fprintf(w, "%s\n", json)
	})
	log.Fatal(http.ListenAndServe(":8081", nil))
}
