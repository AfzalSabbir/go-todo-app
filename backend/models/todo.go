package models

import "time"

// Todo represents a task in the todo list
type Todo struct {
	ID        uint      `gorm:"primaryKey"`
	Title     string    `gorm:"not null"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
}
