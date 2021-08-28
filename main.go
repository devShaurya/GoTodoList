package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func authMiddleware(h http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		bearerAndToken :=  strings.Split(r.Header.Get("Authorization")," ")

		if len(bearerAndToken) != 2 || bearerAndToken[0] != "Bearer" || bearerAndToken[1] != "valid-key" {
            w.Header().Set("Content-Type","application/json")
			w.WriteHeader(http.StatusForbidden)
			json.NewEncoder(w).Encode(map[string]interface{}{"message":"Missing auth token"})
            return
        }
		log.Println("User logged successfully")
		h.ServeHTTP(w,r)	
    })
}

func initializeRouter() *mux.Router{
	
	r := mux.NewRouter()
    r.Use(authMiddleware)
	r.HandleFunc("/todo",func(rw http.ResponseWriter, r *http.Request) {
		rw.Header().Set("Content-Type","application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(map[string]interface{}{"message":"Hi"})
	}).Methods("GET")


	return r
}

func main(){

	r := initializeRouter()
	log.Println("Starting Server at port 5000")
	log.Fatal(http.ListenAndServe(":5000",r))
}