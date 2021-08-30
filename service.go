package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// func validateRequest(todoItem TodoItem) []string {
// 	errs := []string{}

// 	if todoItem.Name == "" || todoItem.Priority == "" || !todoItem.DueDate.Valid {
// 		errs = append(errs, fmt.Errorf("name, priority and dueDate are required").Error())
// 	}

// 	if len(todoItem.Name) > 256{
// 		errs = append(errs, fmt.Errorf("name can only be of 256").Error() )
// 	}

// 	if todoItem.Completed && !todoItem.CompletionDate.Valid {
// 		errs = append(errs, fmt.Errorf("give completionDate if todo item is completed").Error())
// 	}

// 	if todoItem.Priority != "HIGH" && todoItem.Priority != "MED" && todoItem.Priority != "LOW" {
// 		errs = append(errs, fmt.Errorf("priority can be only LOW, MED or HIGH").Error())
// 	}

// 	return errs
// }

func handleDatabaseErrors(err string) string{
	err = strings.Trim(err," ")
	if strings.Contains(err,"1048") {  // mysql code error for null
		return "name, priority and dueDate should be given."

	}else if strings.Contains(err,"1406"){ // mysql code error for size limit crossed
		return "name can be of 256 chars only."

	}else if strings.Contains(err,"3819"){ // mysql code error for violating check constraint
		
		if strings.Contains(err,"priority_check"){
			return "priority can only be 'LOW', 'MED', 'HIGH'."

		}else if strings.Contains(err, "name_length_check"){
			return "name should be given and should be atleast 4 chars and at max 256 chars"

		}else if strings.Contains(err, "completion_date_check"){
			return "if todo item has been completed then completionDate is necessary"
		}
	}else if err == "record not found" { // mysql message for no record
 		return "Todo Item with that id is not present"
	}
	return err
} 

func createTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Post request")
	var todoItem TodoItem
	json.NewDecoder(r.Body).Decode(&todoItem)
	// errs := validateRequest(todoItem)

	todoItem.Name = strings.Trim(todoItem.Name," ")
	rw.Header().Set("Content-Type","application/json")
	// if len(errs) != 0 {
	// 	rw.WriteHeader(http.StatusBadRequest)
	// 	json.NewEncoder(rw).Encode(map[string]interface{}{"messge":errs})
	// 	return
	// }

	if result := db.Create(&todoItem); result.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}
	

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Successfully Created with ID = %d",todoItem.ID)})
	log.Println("Post request successfully executed")
}

func getAllTodoItems (rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type","application/json")
	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"In get all"})
}

func getTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Get request")
	rw.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer id"})
		return
	}

	var todoItem TodoItem
	if result := db.First(&todoItem,params["id"]); result.Error != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	} 

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Success","todoItem":todoItem})
	log.Println("Get request successfully executed")
}

func updateTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Put request")
	var putRequest TodoItem
	json.NewDecoder(r.Body).Decode(&putRequest)
	
	rw.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer id"})
		return
	}

	var todoItem TodoItem
	if result := db.First(&todoItem,params["id"]); result.Error != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}
	
	if result := db.Model(&todoItem).Updates(&putRequest); result.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Todo item with id = %s update successfully",params["id"])})
	log.Println("Put request successfully executed")
}
func deleteTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Delete request")
	rw.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer id"})
		return
	}

	result := db.Delete(&TodoItem{},params["id"])

	if result.Error != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Todo Item with id=%s deleted",params["id"])})
	log.Println("Delete request successfully executed")
}