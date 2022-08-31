package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path"
	"path/filepath"
)

type PrintingResponse struct {
	Printers int
}

type PrintingHandler struct {
}

func NewPrintingHandler() *PrintingHandler {
	return &PrintingHandler{}
}

func (m *PrintingHandler) GetPrinting(w http.ResponseWriter, r *http.Request) {

	resp := MetricResponse{
		Printers: 2,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSOn marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}

func StartPrint() {

	fileName := "benchy-2.gcode"
	url := "http://octopi.local/api/files/local"

	fileDir, _ := os.Getwd()
	filePath := path.Join(fileDir, fileName)

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("path", "/")
	writer.WriteField("print", "false")

	writer.Close()

	fmt.Printf("sending post request %s", url)

	r, _ := http.NewRequest("POST", url, body)

	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Api-Key", "xxx")

	client := &http.Client{}
	client.Do(r)

	fmt.Println("")
	fmt.Println("")
}
