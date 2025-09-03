package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"example.com/todoapi/internal/api"
)

type server struct {
	todos []api.Todo
	next  int64
}

// GET /todos
func (s *server) GetTodos(w http.ResponseWriter, r *http.Request) {
	api.RespondWithJSON(w, http.StatusOK, s.todos)
}

// POST /todos
func (s *server) PostTodos(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body api.NewTodo
	if err := api.BindJSON(ctx, r, &body); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	s.next++
	newItem := api.Todo{Id: s.next, Title: body.Title, Done: false}
	s.todos = append(s.todos, newItem)
	api.RespondWithJSON(w, http.StatusCreated, newItem)
}

// GET /todos/{id}
func (s *server) GetTodosId(w http.ResponseWriter, r *http.Request, id int64) {
	for _, todo := range s.todos {
		if todo.Id == id {
			api.RespondWithJSON(w, http.StatusOK, todo)
			return
		}
	}
	api.RespondWithError(w, http.StatusNotFound, "Todo not found")
}

// PUT /todos/{id}
func (s *server) PutTodosId(w http.ResponseWriter, r *http.Request, id int64) {
	ctx := r.Context()
	var body api.Todo
	if err := api.BindJSON(ctx, r, &body); err != nil {
		api.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	
	for i, todo := range s.todos {
		if todo.Id == id {
			body.Id = id // IDは変更させない
			s.todos[i] = body
			api.RespondWithJSON(w, http.StatusOK, body)
			return
		}
	}
	api.RespondWithError(w, http.StatusNotFound, "Todo not found")
}

// DELETE /todos/{id}
func (s *server) DeleteTodosId(w http.ResponseWriter, r *http.Request, id int64) {
	for i, todo := range s.todos {
		if todo.Id == id {
			// スライスから削除
			s.todos = append(s.todos[:i], s.todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	api.RespondWithError(w, http.StatusNotFound, "Todo not found")
}

func main() {
	s := &server{todos: []api.Todo{}, next: 0}

	r := chi.NewRouter()

	// リクエスト検証（OpenAPIに沿っているか）をミドルウェアで行いたい場合は以下も有用
	// swagger, _ := api.GetSwagger()
	// r.Use(api.OapiRequestValidator(swagger))

	// 生成されたルーターに自分の実装を差し込む
	apiHandler := api.NewStrictHandler(s, nil)
	api.RegisterHandlers(r, apiHandler)

	log.Println("listening on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}
}

// --- 便利ユーティリティ（最小実装用） ---

// api.BindJSON / api.RespondWithJSON / api.RespondWithError は
// oapi-codegenのテンプレに入ってないことがあるので補助関数が必要な場合は
// 自前で用意してください。簡単版を下に置いておきます。
