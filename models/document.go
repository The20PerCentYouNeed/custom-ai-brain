package models

import (
	"time"

	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	tokenizer "github.com/The20PerCentYouNeed/custom-ai-brain/services/tokenize"
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

func (d *Document) GenerateChunks() error {
	chunks, err := tokenizer.ChunkText(d.Content, 10, 3)
	if err != nil {
		return err
	}

	vectors, err := openai.GenerateEmbeddings(chunks)
	if err != nil {
		return err
	}

	documentChunks := make([]Chunk, len(chunks))
	for i, chunk := range chunks {
		documentChunks[i] = Chunk{
			Content:   chunk,
			Embedding: vectors[i],
		}
	}

	d.Chunks = documentChunks
	return nil
}
