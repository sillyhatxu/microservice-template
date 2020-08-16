package model

import "time"

// Model Struct
type User struct {
	Id               int64
	LoginName        string
	Password         string
	UserName         string
	Status           bool
	Platform         string
	Age              *int
	Amount           *float64
	Description      *string
	Birthday         *time.Time
	CreatedTime      time.Time
	LastModifiedTime time.Time
}
