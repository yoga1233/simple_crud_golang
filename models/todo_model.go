package models

import (
	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title    string `json:"title"`
	Desc     string `json:"desc"`
	Deadline string `json:"deadline"`
	Status   string `json:"status"`
	UserID   uint   `json:"user_id"`
}
