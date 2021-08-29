package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func validateRequest(todoItem TodoItem) []string {
	errs := []string{}

	if todoItem.Name == "" || todoItem.Priority == "" || !todoItem.DueDate.Valid {
		errs = append(errs, fmt.Errorf("name, priority and dueDate are required").Error())
	}
	if todoItem.Completed && !todoItem.CompletionDate.Valid {
		errs = append(errs, fmt.Errorf("give completionDate if todo item is completed").Error())
	}

	if todoItem.Priority != "HIGH" && todoItem.Priority != "MED" && todoItem.Priority != "LOW" {
		errs = append(errs, fmt.Errorf("priority can be only LOW, MED or HIGH").Error())
	}

	return errs
}

func createTodoItem (rw http.ResponseWriter, r *http.Request) {
	var todoItem TodoItem
	json.NewDecoder(r.Body).Decode(&todoItem)
	errs := validateRequest(todoItem)

	rw.Header().Set("Content-Type","application/json")
	if len(errs) != 0 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":errs})
		return
	}

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
	var todoItem TodoItem
	json.NewDecoder(r.Body).Decode(&todoItem)
	errs := validateRequest(todoItem)

	rw.Header().Set("Content-Type","application/json")
	if len(errs) != 0 {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":errs})
		return
	}

	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi update"})
}
func deleteTodoItem (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi delete"})
}