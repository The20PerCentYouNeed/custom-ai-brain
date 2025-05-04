package models

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Chunk struct {
	ID         uint     `json:"id" gorm:"primaryKey"`
	DocumentID uint     `json:"document_id" gorm:"index;not null"`
	Document   Document `json:"document" gorm:"constraint:OnDelete:CASCADE;"`

	Content string `json:"content" gorm:"type:text;not null;"`

	Embedding pgvector.Vector `json:"embedding" gorm:"type:vector(1536);not null"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
