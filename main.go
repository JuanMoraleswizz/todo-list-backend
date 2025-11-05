package main

import (
	"fmt"
	"log"
	"net/http"
	"todo-list/internal/handler"
	"todo-list/internal/repository"
	"todo-list/internal/usecase"
	"todo-list/pkg/database"

	"github.com/gorilla/mux"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Access-Control-Max-Age", "86400")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Database connection
	db := database.NewPostgresDB()

	// Dependency injection
	todoRepo := repository.NewTodoRepository(db)
	todoUseCase := usecase.NewTodoUseCase(todoRepo)
	todoHandler := handler.NewTodoHandler(todoUseCase)

	// Router setup
	r := mux.NewRouter()
	r.Use(corsMiddleware)

	r.HandleFunc("/todos", todoHandler.GetTodos).Methods("GET")
	r.HandleFunc("/todos", todoHandler.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoHandler.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoHandler.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoHandler.DeleteTodo).Methods("DELETE")
	r.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNoContent)
	})

	port := ":8080"
	fmt.Printf("Server starting on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, r))
}