package models

import (
	"strings"
	"time"

	"github.com/The20PerCentYouNeed/custom-ai-brain/services/openai"
	tokenizer "github.com/The20PerCentYouNeed/custom-ai-brain/services/tokenizers"
)

type Document struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Title string `json:"title"`

	FileID *uint `json:"file_id" gorm:"index"`
	File   *File `json:"file" gorm:"constraint:OnDelete:CASCADE;"`

	Content string `json:"content" gorm:"type:text;not null;"`

	Chunks    []Chunk   `json:"chunks" gorm:"foreignKey:DocumentID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (document *Document) Chunk(chunkSize int, overlap int) error {

	tokens, err := tokenizer.Tokenize(document.Content)

	if err != nil {
		return err
	}

	var chunks []string

	for start := 0; start < len(tokens); start += chunkSize - overlap {
		end := min(start+chunkSize, len(tokens))

		chunkText := strings.ReplaceAll(strings.Join(tokens[start:end], ""), "\u2581", " ")
		chunks = append(chunks, chunkText)

		if end == len(tokens) {
			break
		}
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

	document.Chunks = documentChunks

	return nil
}
