package models

import (
	"time"
)

type Document struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FileID    uint      `gorm:"index"`
	File      File      `gorm:"constraint:OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Chunks    []Chunk   `json:"chunks" gorm:"foreignKey:DocumentID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
