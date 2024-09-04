package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var message string

type requestBody struct {
	Message string `json:"message"`
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	var reqBody requestBody
	json.NewDecoder(r.Body).Decode(&reqBody)
	message = reqBody.Message
}

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, "+message)
}

func main() {
	router := mux.NewRouter()
	// наше приложение будет слушать запросы на localhost:8080/api/hello
	router.HandleFunc("/api/hello", GetHandler).Methods("GET")
	router.HandleFunc("/api/hello", PostHandler).Methods("POST")
	http.ListenAndServe(":8080", router)
}
