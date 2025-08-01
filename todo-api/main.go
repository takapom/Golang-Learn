package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/yourname/todo-api/handler"
	"github.com/yourname/todo-api/internal"
)

func main() {
	router := chi.NewRouter()

	h := &handler.TodoHandler{}
	apiHandler := internal.HandlerFromMux(h, router)

	log.Println("Starting server on :8080")
	log.Fatal(http.ListenAndServe(":8080", apiHandler))
}
