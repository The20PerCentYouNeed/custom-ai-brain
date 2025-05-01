package models

import (
	"time"

	"github.com/pgvector/pgvector-go"
)

type Chunk struct {
	ID          uint            `gorm:"primaryKey"`
	DocumentID  uint            `gorm:"index;not null"`
	Document    Document        `gorm:"constraint:OnDelete:CASCADE;"`
	Title       string          `json:"title" gorm:"type:text;not null;collate:utf8mb4_unicode_ci"`
	Embedding   pgvector.Vector `gorm:"type:vector(1536);not null"`
	StartOffset int             `gorm:"not null"`
	EndOffset   int             `gorm:"not null"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
