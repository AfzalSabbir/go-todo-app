package models

import (
	"time"
)

// Status represents the status of a Todo (enum-like behavior)
type Status string

const (
	Pending  Status = "pending"
	Complete Status = "complete"
)

// Todo represents a task in the todo list
type Todo struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Task      string    `gorm:"not null" json:"task"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	Status    Status    `gorm:"type:text;check:status IN ('pending', 'complete');default:'pending'" json:"status"`
}
