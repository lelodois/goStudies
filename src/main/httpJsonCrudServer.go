package main

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"strconv"
)

var users = make(map[int]user)

type user struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	LastName   string `json:"lastName"`
	Profession struct {
		Name       string `json:"name"`
		Experience int    `json:"experience"`
	} `json:"profession"`
}

func main() {
	router := httprouter.New()
	router.POST("/users", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		add(writer, request)
	})

	router.GET("/users", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		response(writer, users, http.StatusOK)
	})

	router.GET("/users/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		number, _ := strconv.ParseInt(params.ByName("id"), 10, 32)
		u := users[int(number)]
		if u.Id == 0 {
			response(writer, nil, http.StatusNotFound)
		} else {
			response(writer, u, http.StatusOK)
		}
	})

	router.DELETE("/users/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		number, _ := strconv.ParseInt(params.ByName("id"), 10, 32)
		delete(users, int(number))
		response(writer, user{}, http.StatusOK)
	})

	router.PUT("/users/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		number, _ := strconv.ParseInt(params.ByName("id"), 10, 32)
		update(writer, request, int(number))
	})
	log.Fatal(http.ListenAndServe(":8080", router))
}

func update(writer http.ResponseWriter, i *http.Request, id int) {
	var newU user
	json.NewDecoder(i.Body).Decode(&newU)
	newU.Id = id
	users[id] = newU
	response(writer, newU, http.StatusOK)
}

func add(writer http.ResponseWriter, i *http.Request) {
	var u user
	json.NewDecoder(i.Body).Decode(&u)

	u.Id = len(users) + 1
	users[u.Id] = u

	response(writer, u, http.StatusCreated)
}

func response(writer http.ResponseWriter, u interface{}, httpStatus int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(httpStatus)
	if u != nil {
		json, _ := json.Marshal(u)
		fmt.Fprintf(writer, "%s\n", json)
	}
}
