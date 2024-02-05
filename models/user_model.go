package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `json:"email" validate:"required" gorm:"unique"`
	Password string `json:"password"`
}
