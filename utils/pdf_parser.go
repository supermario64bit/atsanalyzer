package utils

import (
	"bytes"
	"io/ioutil"
	"mime/multipart"

	"rsc.io/pdf"
)

func ExtractTextFromPDF(pdfFile *multipart.FileHeader) (string, error) {
	f, err := pdfFile.Open()
	if err != nil {
		return "", err
	}
	defer f.Close()

	pdfBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return "", err
	}

	reader := bytes.NewReader(pdfBytes)
	r, err := pdf.NewReader(reader, int64(len(pdfBytes)))
	if err != nil {
		return "", err
	}

	var text string
	for i := 1; i <= r.NumPage(); i++ {
		p := r.Page(i)
		if p.V.IsNull() {
			continue
		}
		content := p.Content()
		for _, txt := range content.Text {
			text += txt.S
			text += " "
		}
	}

	return text, nil

}
