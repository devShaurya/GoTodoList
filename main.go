package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var dbErr error

// initialize and automigrate todoItems table
func databaseMigrations(){
	DSN := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC",C.Database.DbUser,C.Database.DbPassword,C.Database.DbBaseUrl,C.Database.DbPort,C.Database.DbName)
	db, dbErr = gorm.Open(mysql.Open(DSN), &gorm.Config{})
	if dbErr!= nil{
		log.Fatalf("Connection to database failed because %s",dbErr.Error())
	}
	if err:= db.AutoMigrate(&TodoItem{}); err != nil{
		log.Fatalf("AutoMigration failed because %s",err.Error())
	}
	log.Println("Database Connected...")
}
// middleware for verifying api requests
func authMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerAndToken :=  strings.Split(r.Header.Get("Authorization")," ")

		if len(bearerAndToken) != 2 || bearerAndToken[0] != "Bearer" || bearerAndToken[1] != C.AuthToken {
            w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"message":"Missing auth token"})
            return
        }
		log.Println("user verified successfully")
		h.ServeHTTP(w,r)	
    })
}
// initialize router
func initializeRouter() *mux.Router{
	
	r := mux.NewRouter()
    r.Use(authMiddleware)
	r.HandleFunc("/todo", getAllTodoItems).Methods("GET")
	r.HandleFunc("/todo/{id}", getTodoItem).Methods("GET")
	r.HandleFunc("/todo", createTodoItem).Methods("POST")
	r.HandleFunc("/todo/{id}", updateTodoItem).Methods("PUT")
	r.HandleFunc("/todo/{id}", deleteTodoItem).Methods("DELETE")
	
	return r
}

func main(){
	loadConfigurations()
	databaseMigrations()
	r := initializeRouter()
	log.Printf("Starting Server at port %d\n",C.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d",C.Server.Port),r))
}