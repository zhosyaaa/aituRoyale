package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	username string `json:"username"`
	email    string `json:"email"`
	password string `json:"password"`
}
