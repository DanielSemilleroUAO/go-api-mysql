package main

import (
	"apiGoSQL/handlers"
	// "apiGoSQL/models"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// models.MigrarUser()

	mux := mux.NewRouter()

	mux.HandleFunc("/api/users",handlers.GetUsers).Methods("GET")
	mux.HandleFunc("/api/users/{id:[0-9]+}",handlers.GetUser).Methods("GET")
	mux.HandleFunc("/api/users",handlers.CreateUser).Methods("POST")
	mux.HandleFunc("/api/users/{id:[0-9]+}",handlers.UpdateUser).Methods("PUT")
	mux.HandleFunc("/api/users/{id:[0-9]+}",handlers.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", mux))

}
