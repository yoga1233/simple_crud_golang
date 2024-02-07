package models

import (
	"gorm.io/gorm"
	"time"
)

type Todo struct {
	gorm.Model
	Title    string    `json:"title"`
	Desc     string    `json:"desc"`
	Deadline time.Time `json:"deadline"`
	UserID   uint      `json:"user_id"`
}
