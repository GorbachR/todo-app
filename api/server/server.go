package server

import (
	"encoding/json"
	"github.com/GorbachR/todo-app/api/data/query"
	"github.com/GorbachR/todo-app/api/data/schema"
	"github.com/GorbachR/todo-app/api/utility"
	"net/http"
	"strconv"
)

type dataId struct {
	Id int `json:"id"`
}

func SetupServer(listenAddr string) *http.Server {
	mux := http.NewServeMux()

	mux.HandleFunc("/todos", handleTodos)
	mux.HandleFunc("/todos/", handleTodo)

	return &http.Server{
		Addr:    "127.0.0.1" + listenAddr,
		Handler: mux,
	}
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		getTodos(w, r)
	case http.MethodPost:
		postTodo(w, r)
	default:
		http.Error(w, "Method not implemented", http.StatusMethodNotAllowed)
	}
}

func handleTodo(w http.ResponseWriter, r *http.Request) {

	// Implementation to match /todos/id
	todoId, err := utility.ExtractIdUrl(r.URL.Path, 2)

	if err != nil {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		getTodo(w, r, todoId)
	case http.MethodPut:
		putTodo(w, r, todoId)
	case http.MethodDelete:
		deleteTodo(w, r, todoId)
	default:
		http.Error(w, "Method not implemented", http.StatusMethodNotAllowed)
	}
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()

	limit, offset := 5, 0

	if _, exists := queryParams["offset"]; exists {
		var err error
		offset, err = strconv.Atoi(queryParams.Get("offset"))
		if err != nil {
			http.Error(w, "Invalid offset param", http.StatusUnprocessableEntity)
			return
		}
	}

	if _, exists := queryParams["limit"]; exists {
		var err error
		offset, err = strconv.Atoi(queryParams.Get("limit"))
		if err != nil {
			http.Error(w, "Invalid limit param", http.StatusUnprocessableEntity)
			return
		}
	}

	todos, err := query.GetTodos(limit, offset)

	if err == nil {
		json.NewEncoder(w).Encode(todos)
	}
}

func postTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo schema.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid payload", http.StatusUnprocessableEntity)
		return
	}

	res, err := query.CreateTodo(newTodo)
	resId, resErr := res.LastInsertId()

	if err != nil || resErr != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(dataId{Id: int(resId)})
	w.WriteHeader(http.StatusCreated)
}

func putTodo(w http.ResponseWriter, r *http.Request, todoId int) {
	var newTodo schema.Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		http.Error(w, "Invalid payload", http.StatusUnprocessableEntity)
		return
	}

	err := query.UpdateTodo(newTodo)

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func deleteTodo(w http.ResponseWriter, r *http.Request, todoId int) {

	err := query.DeleteTodo(todoId)

	if err != nil {
		http.Error(w, "Server error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func getTodo(w http.ResponseWriter, r *http.Request, todoId int) {
	row, err := query.GetTodo(todoId)

	if err != nil {
		http.Error(w, "Server Error", http.StatusInternalServerError)
	}

	json.NewEncoder(w).Encode(&row)
}
