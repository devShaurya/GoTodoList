package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"gorm.io/gorm"
)

// custom data type to store dates i.e. DueDate and CompletionDate
// format --> dd/mm/YYYY
type Date struct{
	sql.NullTime
}

// marshal function for custom data type
func (date Date) MarshalJSON() ([]byte, error) {
	if(date.Valid){
		return json.Marshal(fmt.Sprintf("%02d-%02d-%d",date.Time.Day(),date.Time.Month(),date.Time.Year()))
	}
	return json.Marshal(nil)
}

// unmarshal function for custom data type
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

// model for mysql database
// constraints involve --> 
// 1. Name, Priority and DueDate should not be null.
// 2. String size of Name should be atleast 4 chars and atmost 256 chars.
// 3. Priority can only be one of LOW, MED and HIGH
// 4. If TodoItem is completed (i.e. Completed is true) then Completion date should be also be present and similarly 
// if it isn't completed (i.e. Completed is false) then CompletionDate can't be present
type TodoItem struct{
	gorm.Model
	Name 			string 		`json:"name" gorm:"not null;size:256;check:name_length_check,length(name)>3"`
	Description		string 		`json:"description"`
	Priority		string      `json:"priority" gorm:"not null;default:'LOW';check:priority_check,priority in ('HIGH', 'LOW', 'MED')"`
	DueDate			Date        `json:"dueDate"  gorm:"not null"`
	Completed		bool   		`json:"completed" gorm:"default:false"`
	CompletionDate 	Date 	    `json:"completionDate" gorm:"check:completion_date_check,(completed = false AND completion_date IS NULL) OR (completed = true AND completion_date IS NOT NULL) "`
}
