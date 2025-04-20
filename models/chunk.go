package models

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Chunk struct {
	ID         uint            `gorm:"primaryKey"`
	DocumentID uint            `gorm:"index;not null"`
	Document   Document        `gorm:"constraint:OnDelete:CASCADE;"`
	Content    string          `gorm:"type:text;not null"`
	Embedding  pgvector.Vector `gorm:"type:vector(1536);not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
