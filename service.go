package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

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
 		return "Todo Item with that ID is not present"
	
	} else if strings.Contains(err, "1054"){  // mysql error #1054 - Unknown column in 'Field List'
		return "name, description, priority, dueDate, completed and completionDate are the only fields that todo item contains"
	}
	return err
} 

// Create new todo item
// name, priority and dueDate should be present in request. 
// dueDate and completionDate must follow dd-mm-yyyy format.
// name should be atleast 4 chars and atmost 256 chars.
// priority should be one of LOW, MED or HIGH.
// Default value of completed is false and priority is LOW.
// if completed is true, then completionDate should also be provided.
func createTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Post request")
	var todoItem TodoItem
	json.NewDecoder(r.Body).Decode(&todoItem)
	
	todoItem.Name = strings.Trim(todoItem.Name," ")
	rw.Header().Set("Content-Type","application/json")
	
	if result := db.Create(&todoItem); result.Error != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}
	

	rw.WriteHeader(http.StatusCreated)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Successfully Created with ID = %d",todoItem.ID)})
	log.Println("Post request successfully executed")
}

// Get All Todo List items
// Filtering can be done using parameters: priority, name, description and completed.
// Multiple parameters are allowed.
// strictly follow names of parameters for filtering.
// if string mentioned in name or description is contained in respective fields then they are returned.
// string mentioned in priority is searched for exact match.
// completed can be true or false. If anything other than true is mentioned, its considered as false.
func getAllTodoItems (rw http.ResponseWriter, r *http.Request) {
	log.Println("get all request")
	rw.Header().Set("Content-Type","application/json")
	var todoItems []TodoItem
	
	query, namedArgument := getQueryString(r.URL.Query())

	
	if result := db.Raw(query,namedArgument).Scan(&todoItems); result.Error != nil{
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":fmt.Sprintf("Error %s",result.Error.Error())})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Success","todoItems":todoItems})
	log.Println("get all request successfully executed")
}

// get todo item according to ID present in request path
func getTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Get request")
	rw.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer ID"})
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

// Updates todo item according to ID present in request path
// first checks if an item with requested ID is present or not. 
// single or multiple fields can be updated.
// but they must follow constraints present in database.
func updateTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Put request")
	var putRequest TodoItem
	json.NewDecoder(r.Body).Decode(&putRequest)
	
	rw.Header().Set("Content-Type","application/json")
	
	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer ID"})
		return
	}

	// check if item with requested ID is present
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
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Todo item with ID = %s update successfully",params["id"])})
	log.Println("Put request successfully executed")
}

// deletes a todo item according to ID present in request path
// actually it soft deletes it. Meaning it does not remove row from table but updates deleted_at column
// with the time when request was made.
func deleteTodoItem (rw http.ResponseWriter, r *http.Request) {
	log.Println("Delete request")
	rw.Header().Set("Content-Type","application/json")

	params := mux.Vars(r)
	if _, err := strconv.Atoi(params["id"]); err != nil {
		rw.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge":"Give integer ID"})
		return
	}
	// check if item with requested ID has already been deleted
	var todoItem TodoItem
	if result := db.First(&todoItem,params["id"]); result.Error != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}
	result := db.Delete(&TodoItem{},params["id"])

	if result.Error != nil{
		rw.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(rw).Encode(map[string]interface{}{"messge": handleDatabaseErrors(result.Error.Error())})
		return
	}

	rw.WriteHeader(http.StatusOK)
	json.NewEncoder(rw).Encode(map[string]interface{}{"message":fmt.Sprintf("Todo Item with ID=%s deleted",params["id"])})
	log.Println("Delete request successfully executed")
}

// used in sql query, for replacing named arguments in the query
type NamedArgument struct{
	Priority string
	Completed bool
	Name string
	Description string
}

// To make sql query to be used for filtering the todo items
func getQueryString(q url.Values) (string, NamedArgument) {
	var namedArgument = NamedArgument{}
	query := "SELECT * FROM todo_items WHERE deleted_at IS NULL"

	if q.Has("priority"){
		query += " AND priority = @Priority"
		namedArgument.Priority = q.Get("priority")
	}
	if q.Has("completed"){
		query += " AND completed = @Completed"
		if q.Get("completed")=="true" {
			namedArgument.Completed = true
		} else{
			namedArgument.Completed = false
		}
	}
	if q.Has("name"){
		query += " AND name LIKE @Name"
		namedArgument.Name = "%"+q.Get("name")+"%"
	}
	if q.Has("description"){
		query += " AND description LIKE @Description"
		namedArgument.Description = "%"+q.Get("description")+"%"
	}

	return query, namedArgument
}