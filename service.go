package main

import (
	"encoding/json"
	"net/http"
)

func createTodoItem (rw http.ResponseWriter, r *http.Request) {
	var todoItem TodoItemCreateRequest
	json.NewDecoder(r.Body).Decode(&todoItem)
	

	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"item":todoItem,"message":"Successfully Created"})
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