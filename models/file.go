package models

import (
	"os"
	"os/exec"
	"regexp"
	"strings"
	"time"

	"github.com/The20PerCentYouNeed/custom-ai-brain/utils"
)

type File struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Type      string    `json:"type"`
	Source    string    `json:"source"`
	Uri       string    `json:"uri"`
	MimeType  string    `json:"mime_type"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (f *File) ExtractTextFromPDF() (string, error) {
	filePath := utils.StoragePath("files", f.Uri)

	cmd := exec.Command("pdftotext", "-layout", filePath, "-")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	cleanedText := sanitizeText(string(output))

	return cleanedText, nil
}

func (f *File) ExtractTextFromTXT() (string, error) {
	filePath := utils.StoragePath("files", f.Uri)

	content, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	cleanedText := sanitizeText(string(content))

	return cleanedText, nil
}

func (f *File) ExtractTextFromDOCX() (string, error) {
	filePath := utils.StoragePath("files", f.Uri)

	cmd := exec.Command("pandoc", filePath, "-t", "plain")
	output, err := cmd.Output()
	if err != nil {
		return "", err
	}

	cleanedText := sanitizeText(string(output))

	return cleanedText, nil
}

func sanitizeText(input string) string {
	input = strings.ReplaceAll(input, "\t", " ")

	spaceRe := regexp.MustCompile(`\s+`)
	input = spaceRe.ReplaceAllString(input, " ")

	lines := strings.Split(input, "\n")
	cleanedLines := make([]string, 0, len(lines))

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if len(line) > 0 {
			cleanedLines = append(cleanedLines, line)
		}
	}

	return strings.Join(cleanedLines, "\n")
}
