package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AankTia/go-simple-crud/database"
	"github.com/AankTia/go-simple-crud/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// Connetct to database
	database.ConnectDB()

	// Initialize router
	router := mux.NewRouter()

	// Routes
	router.HandleFunc("/api/tasks", handlers.GetAllTasks).Methods("GET")
	router.HandleFunc("/api/tasks/{id}", handlers.GetTask).Methods("GET")
	router.HandleFunc("/api/tasks", handlers.CreateTask).Methods("POST")
	router.HandleFunc("/api/tasks/{id}", handlers.UpdateTask).Methods("PUT")
	router.HandleFunc("/api/tasks/{id}", handlers.DeleteTask).Methods("DELETE")

	// Serve static files
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static")))

	// Start server
	port := ":8080"
	fmt.Printf("Server is running on htttp://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}