package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Server() {
	router := mux.NewRouter()
	router.HandleFunc("/", handleHome).Methods("GET")
}

func main() {
	router := Server()
	log.Fatal(http.ListenAndServe(":8080", router))

}

func handleHome(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(map[string]interface{}{
		"status":  200,
		"message": "Hello World",
	})
	log.Println("** Service Started on Port 8080 **")
	fmt.Println("** Service Started on Port 8080 **")
}
