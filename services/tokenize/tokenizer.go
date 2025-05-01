package tokenizer

import (
	"fmt"

	tiktoken "github.com/pkoukk/tiktoken-go"
)

type ChunkData struct {
	Text        string
	StartOffset int
	EndOffset   int
}

func ChunkText(text string, chunkSize int, overlap int) ([]ChunkData, error) {
	enc, err := tiktoken.EncodingForModel("text-embedding-3-small")
	if err != nil {
		return nil, err
	}

	tokens := enc.Encode(text, nil, nil)

	var chunks []ChunkData

	var count int

	for start := 0; start < len(tokens); start += chunkSize - overlap {
		end := min(start+chunkSize, len(tokens))

		chunkText := enc.Decode(tokens[start:end])

		// Find the byte offsets
		chunk := string(chunkText)

		startByteOffset := len(enc.Decode(tokens[:start]))
		endByteOffset := len(enc.Decode(tokens[:end]))

		chunks = append(chunks, ChunkData{
			Text:        chunk,
			StartOffset: startByteOffset,
			EndOffset:   endByteOffset,
		})

		if end == len(tokens) {
			break
		}

		count++
		fmt.Println("Completed Chunks", count)
	}

	fmt.Println("Total Chunks", count)
	return chunks, nil
}
