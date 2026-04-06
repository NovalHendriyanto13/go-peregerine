package libraries

import (
	"github.com/google/uuid"
	"math/rand"
	"bytes"
	"github.com/ledongthuc/pdf"
)

func GenerateUUID() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}

func RandomString(n int) (string) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func ReadPdfToText(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}

	defer f.Close()

	var buf bytes.Buffer
	totalPage := r.NumPage()
	
	for i := 1; i <= totalPage; i++ {
		page := r.Page(i)
		if page.V.IsNull() {
			continue
		}

		content, _ := page.GetPlainText(nil)
		buf.WriteString(content)
	}

	return buf.String(), nil
}