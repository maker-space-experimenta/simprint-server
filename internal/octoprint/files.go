package octoprint

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

// Sends a file to a octoprint instance by post http
func (m *Octoprinter) postFiles(file *os.File, context context.Context, print bool) error {
	m.logger.Infof("running PostFiles for %v ", m.hostname)

	if file == nil {
		m.logger.Errorf("file is nil")
		return errors.New("file is nil")
	}

	addr, err := net.LookupIP(m.hostname)
	if err != nil {
		m.logger.Errorf("Unknown host " + m.hostname)
		return err
	} else {
		m.logger.Debugf("IP address: %v", addr)
	}

	urlApi := fmt.Sprintf("%v://%v/api", "http", addr[0])
	m.logger.Debugf(urlApi)

	urlFiles, err := url.JoinPath(urlApi, "files", "local")
	if err != nil {
		m.logger.Errorf("FATAL: error on joining paths, %v", err)
		return err
	}
	m.logger.Debugf(urlFiles)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("path", "/")
	writer.WriteField("print", fmt.Sprintf("%v", print))
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

func (m *Octoprinter) SendFile(file *os.File, context context.Context) error {
	return m.postFiles(file, context, false)
}

func (m *Octoprinter) PrintFile(file *os.File, context context.Context) error {
	return m.postFiles(file, context, true)
}
