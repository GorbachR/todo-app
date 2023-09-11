package server

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/GorbachR/todo-app/api/data/query"
)

func SetupServer(listenAddr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", handleTodos)

	return &http.Server{
		Addr:    "127.0.0.1" + listenAddr,
		Handler: mux,
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	offset := 0

	if _, exists := queryParams["offset"]; exists {
		offset, _ = strconv.Atoi(queryParams.Get("offset"))
	}

	todos, err := query.Todos(offset)

	if err == nil {
		json.NewEncoder(w).Encode(todos)
	}
}
