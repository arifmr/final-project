package model

import (
	"time"

	"gorm.io/gorm"
)

type Todo struct {
	gorm.Model
	Title       string    `gorm:"column:title"`
	Status      string    `gorm:"column:status"`
	Description string    `gorm:"column:description"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	Deadline    time.Time `gorm:"column:deadline"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type ErrorResponse struct {
	Message string `json:"message"`
}
