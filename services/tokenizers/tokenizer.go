package tokenizer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os/exec"
	"unicode/utf8"
)

func Tokenize(text string) ([]string, error) {

	if !utf8.ValidString(text) {
		return nil, fmt.Errorf("input text contains invalid UTF-8")
	}

	cmd := exec.Command("python3", "scripts/tokenizer.py")
	var stdin bytes.Buffer
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdin = &stdin
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	_, err := stdin.WriteString(text)
	if err != nil {
		return nil, err
	}

	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	var tokens []string
	err = json.Unmarshal(stdout.Bytes(), &tokens)
	if err != nil {
		return nil, err
	}

	return tokens, nil
}
