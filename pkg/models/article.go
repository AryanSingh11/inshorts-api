package models

import (
	//"github.com/go-playground/locales/id"
	//"github.com/golang/protobuf/ptypes/timestamp"
	"time"
)

//this function returns the time it was invoked
//in string format
// .Format method returns formatted time in string format
func StrTime() string {
	current := time.Now()
	strTime := current.Format("2006-01-02 15:04:05")
	return strTime
}

type Article struct {
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Subtitle  string `json:"subtitle"`
	Content   string `json:"content"`
	Timestamp string
}
