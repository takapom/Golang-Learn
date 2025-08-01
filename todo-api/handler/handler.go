package handler

import (
	"encoding/json"
	"net/http"

	"github.com/yourname/todo-api/internal"
)

type TodoHandler struct{}

// GET /todos
func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos := []internal.Todo{
		{Id: ptr("1"), Task: ptr("Buy milk"), Done: ptr(false)},
		{Id: ptr("2"), Task: ptr("Do homework"), Done: ptr(true)},
	}
	json.NewEncoder(w).Encode(todos)
}

// POST /todos
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo internal.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	// 仮で ID 付与
	todo.Id = ptr("3")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// GET /todos/{id}
func (h *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request, id string) {
	todo := internal.Todo{Id: &id, Task: ptr("Mock task"), Done: ptr(false)}
	json.NewEncoder(w).Encode(todo)
}

// PUT /todos/{id}
func (h *TodoHandler) UpdateTodo(w http.ResponseWriter, r *http.Request, id string) {
	var todo internal.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	todo.Id = &id
	json.NewEncoder(w).Encode(todo)
}

// DELETE /todos/{id}
func (h *TodoHandler) DeleteTodo(w http.ResponseWriter, r *http.Request, id string) {
	w.WriteHeader(http.StatusNoContent)
}

// 補助関数：値のアドレスを取得
func ptr[T any](v T) *T {
	return &v
}
