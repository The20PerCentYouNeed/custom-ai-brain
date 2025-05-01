package models

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
	path := filepath.Join(utils.StoragePath(), "files", f.Uri)

	outputPath := path + ".txt"

	f.Uri = f.Uri + ".txt"

	cmd := exec.Command("pdftotext", "-layout", "-enc", "UTF-8", path, outputPath)
	err := cmd.Run()
	if err != nil {
		return "", err
	}

	content, err := os.ReadFile(outputPath)
	if err != nil {
		return "", err
	}

	// Clean/sanitize the text
	cleaned := sanitizeText(string(content))

	// Step 4: Overwrite the .txt file with the cleaned content
	err = os.WriteFile(outputPath, []byte(cleaned), 0644)
	if err != nil {
		return "", err
	}

	fmt.Println("cleaned", cleaned)
	return cleaned, nil
}

func sanitizeText(input string) string {
	// 1. Replace all tabs with a single space
	input = strings.ReplaceAll(input, "\t", " ")

	// 2. Replace multiple spaces with a single space
	spaceRe := regexp.MustCompile(`\s+`)
	input = spaceRe.ReplaceAllString(input, " ")

	// 3. Split into lines and remove empty lines
	lines := strings.Split(input, "\n")
	var cleanedLines []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" {
			cleanedLines = append(cleanedLines, line)
		}
	}

	// 4. Join back the cleaned lines with a single newline
	return strings.Join(cleanedLines, "\n")
}

func (f *File) ExtractTextFromTXT() (string, error) {
	// return content.String(), nil
	return "", nil
}

func (f *File) ExtractTextFromDOCX() (string, error) {
	// // Create a temporary file for the DOCX content
	// tmpFile, err := os.CreateTemp("", "docx-*.docx")
	// if err != nil {
	// 	return "", err
	// }
	// defer os.Remove(tmpFile.Name())
	// defer tmpFile.Close()

	// // Write the content to the temporary file
	// if _, err := tmpFile.Write(); err != nil {
	// 	return "", err
	// }

	// doc, err := document.Open(tmpFile.Name())
	// if err != nil {
	// 	return "", err
	// }
	// defer doc.Close()

	// var text bytes.Buffer
	// for _, para := range doc.Paragraphs() {
	// 	for _, run := range para.Runs() {
	// 		text.WriteString(run.Text())
	// 	}
	// 	text.WriteString("\n")
	// }

	// return text.String(), nil
	return "", nil
}
