package openai

import (
	"context"
	"os"

	"github.com/pgvector/pgvector-go"
	openai "github.com/sashabaranov/go-openai"
)

func GenerateEmbeddings(texts []string) ([]pgvector.Vector, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateEmbeddings(context.Background(), openai.EmbeddingRequest{
		Model: openai.AdaEmbeddingV2,
		Input: texts,
	})

	if err != nil {
		return nil, err
	}

	vectors := make([]pgvector.Vector, len(resp.Data))
	for i, embedding := range resp.Data {
		float32Array := make([]float32, len(embedding.Embedding))
		for j, v := range embedding.Embedding {
			float32Array[j] = float32(v)
		}
		vectors[i] = pgvector.NewVector(float32Array)
	}

	return vectors, nil
}

func GenerateEmbedding(text string) (pgvector.Vector, error) {
	vectors, err := GenerateEmbeddings([]string{text})
	if err != nil {
		return pgvector.Vector{}, err
	}
	return vectors[0], nil
}

func GenerateAnswer(question string, content string) (string, error) {
	client := openai.NewClient(os.Getenv("OPENAI_API_KEY"))
	resp, err := client.CreateChatCompletion(context.Background(), openai.ChatCompletionRequest{
		Model: openai.GPT4o,
		Messages: []openai.ChatCompletionMessage{
			{
				Role:    openai.ChatMessageRoleSystem,
				Content: "You are an assistant that answers user questions using the provided business-specific documents. Use the context provided to answer accurately.",
			},
			{
				Role:    openai.ChatMessageRoleUser,
				Content: "Context:\n" + content + "\n\nQuestion:\n" + question,
			},
		},
	})

	if err != nil {
		return "", err
	}

	return resp.Choices[0].Message.Content, nil
}
