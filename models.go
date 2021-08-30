package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
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
	gorm.Model
	Name 			string 		`json:"name" gorm:"not null;size:256;check:name_length_check,length(name)>3"`
	Description		string 		`json:"description"`
	Priority		string      `json:"priority" gorm:"default:'LOW';check:priority_check,priority in ('HIGH', 'LOW', 'MED')"`
	DueDate			Date        `json:"dueDate"  gorm:"not null"`
	Completed		bool   		`json:"completed" gorm:"default:false"`
	CompletionDate 	Date 	    `json:"completionDate" gorm:"check:completion_date_check,(completed = false AND completion_date IS NULL) OR (completed = true AND completion_date IS NOT NULL) "`
}
