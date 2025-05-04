package models

import (
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

	storagePath := utils.StoragePath()
	pdfPath := filepath.Join(storagePath, "files", f.Uri)
	txtPath := pdfPath + ".txt"

	cmd := exec.Command("pdftotext", "-layout", pdfPath, txtPath)
	if err := cmd.Run(); err != nil {
		return "", err
	}

	content, err := os.ReadFile(txtPath)
	if err != nil {
		return "", err
	}

	cleanedText := sanitizeText(string(content))

	return cleanedText, nil
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
