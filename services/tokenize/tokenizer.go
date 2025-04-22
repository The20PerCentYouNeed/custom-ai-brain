package tokenizer

import (
	tiktoken "github.com/pkoukk/tiktoken-go"
)

func ChunkText(text string, chunkSize int, overlap int) ([]string, error) {
	enc, err := tiktoken.EncodingForModel("text-embedding-3-small")
	if err != nil {
		return nil, err
	}

	tokens := enc.Encode(text, nil, nil)

	var chunks []string
	for start := 0; start < len(tokens); start += chunkSize - overlap {
		end := min(start+chunkSize, len(tokens))

		chunk := enc.Decode(tokens[start:end])
		chunks = append(chunks, string(chunk))

		if end == len(tokens) {
			break
		}
	}

	return chunks, nil
}
