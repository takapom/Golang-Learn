package main

import (
	"encoding/json"
	"go_sample/api"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type MyServer struct{}

// 実際のデータを返す処理をここで書く
func (s *MyServer) GetCars(w http.ResponseWriter, r *http.Request) {
	cars := []api.Car{
		{Id: ptr("1"), Model: ptr("Civic")},
		{Id: ptr("2"), Model: ptr("Corolla")},
	}

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(cars)
}

// ヘルパー関数：stringをポインタに変換
func ptr(s string) *string {
	return &s
}

func main() {
	r := chi.NewRouter()
	server := &MyServer{}
	handler := api.HandlerFromMux(server, r)

	log.Println("Server started at :8080")
	http.ListenAndServe(":8080", handler)
}
