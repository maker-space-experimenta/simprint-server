package octoprint

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

func (m *Octoprinter) PostFiles(file *os.File) error {
	m.logger.Infof("running PostFiles for %v ", m.hostname)

	apiUrl := fmt.Sprintf("%v://%v/api", "http", m.hostname)

	urlFiles, err := url.JoinPath(apiUrl, "files", "local")
	if err != nil {
		m.logger.Errorf("FATAL: error on joining paths, %v", err)
		return err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("path", "/")
	writer.WriteField("print", "true")
	writer.Close()

	m.logger.Infof(fmt.Sprintf("sending post request %s", urlFiles))

	r, _ := http.NewRequest("POST", urlFiles, body)
	r = r.WithContext(m.ctx)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Api-Key", m.apiKey)

	client := &http.Client{}
	client.Do(r)

	return nil
}
