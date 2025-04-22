package models

import (
	"bytes"
	"os"
	"time"

	"github.com/ledongthuc/pdf"
	"github.com/unidoc/unioffice/document"
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

func (f *File) ExtractTextFromPDF(content bytes.Buffer) (string, error) {
	reader := bytes.NewReader(content.Bytes())

	pdfReader, err := pdf.NewReader(reader, int64(reader.Len()))
	if err != nil {
		return "", err
	}

	var text string

	numPages := pdfReader.NumPage()
	for pageIndex := 1; pageIndex <= numPages; pageIndex++ {
		page := pdfReader.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}
		pageText, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}
		text += pageText
	}

	return text, nil
}

func (f *File) ExtractTextFromTXT(content bytes.Buffer) (string, error) {
	return content.String(), nil
}

func (f *File) ExtractTextFromDOCX(content bytes.Buffer) (string, error) {
	// Create a temporary file for the DOCX content
	tmpFile, err := os.CreateTemp("", "docx-*.docx")
	if err != nil {
		return "", err
	}
	defer os.Remove(tmpFile.Name())
	defer tmpFile.Close()

	// Write the content to the temporary file
	if _, err := tmpFile.Write(content.Bytes()); err != nil {
		return "", err
	}

	doc, err := document.Open(tmpFile.Name())
	if err != nil {
		return "", err
	}
	defer doc.Close()

	var text bytes.Buffer
	for _, para := range doc.Paragraphs() {
		for _, run := range para.Runs() {
			text.WriteString(run.Text())
		}
		text.WriteString("\n")
	}

	return text.String(), nil
}
