package models

import (
	"os"
	"path/filepath"
	"time"

	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	tokenizer "github.com/The20PerCentYouNeed/custom-ai-brain/services/tokenize"
	"github.com/The20PerCentYouNeed/custom-ai-brain/utils"
)

type Document struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	FileID    uint      `gorm:"index"`
	File      File      `gorm:"constraint:OnDelete:CASCADE;"`
	Title     string    `json:"title"`
	Chunks    []Chunk   `json:"chunks" gorm:"foreignKey:DocumentID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (d *Document) GenerateChunks() error {

	path := filepath.Join(utils.StoragePath(), "files", d.File.Uri)

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	chunks, err := tokenizer.ChunkText(string(content), 500, 50)

	if err != nil {
		return err
	}

	var chunkTexts []string
	for _, chunk := range chunks {
		chunkTexts = append(chunkTexts, chunk.Text)
	}

	vectors, err := openai.GenerateEmbeddings(chunkTexts)
	if err != nil {
		return err
	}

	documentChunks := make([]Chunk, len(chunks))
	for i, chunk := range chunks {
		documentChunks[i] = Chunk{
			StartOffset: chunk.StartOffset,
			EndOffset:   chunk.EndOffset,
			Embedding:   vectors[i],
		}
	}

	d.Chunks = documentChunks
	return nil
}
