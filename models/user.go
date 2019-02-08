package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index" json:"username"`
	Email string `gorm:"unique_index" json:"email"`
	Password string `json:"password"`
}
