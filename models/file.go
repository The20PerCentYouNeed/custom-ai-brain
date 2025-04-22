package models

import (
	"fmt"
	"path/filepath"
	"time"

	"github.com/The20PerCentYouNeed/custom-ai-brain/utils"
	pdf "github.com/ledongthuc/pdf"
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

	file, r, err := pdf.Open(path)
	if err != nil {

		return "", err
	}

	defer file.Close()

	var text string
	totalPage := r.NumPage()
	for pageIndex := 1; pageIndex <= totalPage; pageIndex++ {
		page := r.Page(pageIndex)
		if page.V.IsNull() {
			continue
		}

		content, err := page.GetPlainText(nil)
		if err != nil {
			return "", err
		}

		text += content
	}
	fmt.Println(text)
	// reader := bytes.NewReader(content.Bytes())

	// pdfReader, err := pdf.NewReader(reader, int64(reader.Len()))
	// if err != nil {
	// 	return "", err
	// }

	// var text string

	// numPages := pdfReader.NumPage()
	// for pageIndex := 1; pageIndex <= numPages; pageIndex++ {
	// 	page := pdfReader.Page(pageIndex)
	// 	if page.V.IsNull() {
	// 		continue
	// 	}
	// 	pageText, err := page.GetPlainText(nil)
	// 	if err != nil {
	// 		return "", err
	// 	}
	// 	text += pageText
	// }

	return "", nil
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
