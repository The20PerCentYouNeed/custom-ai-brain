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
