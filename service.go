package main

import (
	"encoding/json"
	"net/http"
)

func createTodoItem (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"In create"})
}

func getAllTodoItems (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"In get all"})
}

func getTodoItem (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi"})
}

func updateTodoItem (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi update"})
}
func deleteTodoItem (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi delete"})
}