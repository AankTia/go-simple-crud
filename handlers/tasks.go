package models

import (
	"time"
)

type Task struct {
	ID 			uint 		`json:"id" gorm:"primayKey"`
	Title 		string 		`json:"title"`
	Description string 		`json:"description"`
	Status 		string 		`json:"status"`
	CreatedAt 	time.Time 	`json:"created_at"`
	UpdatedAt 	time.Time 	`json:"updated_at"`
}