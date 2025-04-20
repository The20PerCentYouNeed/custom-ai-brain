package models

import (
	"time"
)

type Document struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Chunks    []Chunk   `json:"chunks" gorm:"foreignKey:DocumentID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
