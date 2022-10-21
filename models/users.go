package models

import "time"

type Users struct {
	ID            int
	Username      string
	Password      string
	FirstName     string
	LastName      string
	TimeInString  string
	TimeOutString string
	TimeIn        time.Time
	TimeOut       time.Time
}
