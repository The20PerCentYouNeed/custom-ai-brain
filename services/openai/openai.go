package openai

import (
	"context"
	"os"

	"github.com/pgvector/pgvector-go"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateEmbedding(text string) (pgvector.Vector, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: []string{text},
	})

	if err != nil {
		return pgvector.Vector{}, err
	}

	float32Array := make([]float32, len(resp.Data[0].Embedding))
	for i, v := range resp.Data[0].Embedding {
		float32Array[i] = float32(v)
	}

	return pgvector.NewVector(float32Array), nil
}
