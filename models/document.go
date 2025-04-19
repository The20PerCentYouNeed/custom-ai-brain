package models

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Document struct {
	ID        uint            `json:"id" gorm:"primaryKey"`
	Title     string          `json:"title"`
	Content   string          `json:"content"`
	Embedding pgvector.Vector `gorm:"type:vector(1536)" json:"embedding"`
	CreatedAt time.Time       `json:"created_at"`
	UpdatedAt time.Time       `json:"updated_at"`
}
