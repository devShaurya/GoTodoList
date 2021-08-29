package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Date struct{
	sql.NullTime
}
func (date Date) MarshalJSON() ([]byte, error) {
	// s := time.Parse("02-01-2006",date.Time)
	if(date.Valid){
		return json.Marshal(fmt.Sprintf("%02d-%02d-%d",date.Time.Day(),date.Time.Month(),date.Time.Year()))
	}
	return json.Marshal(nil)
}
  
func (date *Date) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	if s!="null" {
		t, err := time.Parse("02-01-2006",s)
		if err != nil {
			return err
		}
		date.Time = t
		date.Valid = true
	}else{
		date.Time = time.Time{}
		date.Valid = false
	}
	return nil
}

type TodoItem struct{
	ID              uint        `json:"id"`
	Name 			string 		`json:"name"`
	Description		string 		`json:"description"`
	Priority		string      `json:"priority"`
	DueDate			Date        `json:"dueDate"`
	Completed		bool   		`json:"completed"`
	CompletionDate 	Date 	    `json:"completionDate"`
}
