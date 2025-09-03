package api

import (
	"context"
	"encoding/json"
	"net/http"
)

func BindJSON(_ context.Context, r *http.Request, v any) error {
	return json.NewDecoder(r.Body).Decode(v)
}
func RespondWithJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(v)
}
func RespondWithError(w http.ResponseWriter, status int, msg string) {
	RespondWithJSON(w, status, map[string]string{"error": msg})
}
