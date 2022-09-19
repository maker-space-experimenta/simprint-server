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
	"path/filepath"

	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
)

type PrintHandler struct {
	config util.Config
}

func NewPrintHandler(config util.Config) *PrintHandler {
	return &PrintHandler{
		config: config,
	}
}

func (m *PrintHandler) PostPrint(w http.ResponseWriter, r *http.Request) {

	log.Println("executing PostPrint Endpoint")

	defer r.Body.Close()
	bodyBytes, err := io.ReadAll(r.Body)
	if err != nil {
		panic(err)
	}

	var jsonRes map[string]interface{}
	_ = json.Unmarshal(bodyBytes, &jsonRes)

	file := jsonRes["file"].(string)
	printer := jsonRes["printer"].(string)

	log.Println(file)
	log.Println(printer)

	log.Println("sending print job")

	dirName := m.config.TempFileDir
	filePath := fmt.Sprintf("%v/%v", dirName, file)

	apiUrl := fmt.Sprintf("%v/%v", printer, "files/local")

	StartPrint(apiUrl, filePath)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{ \"result\": \"ok\" }"))
}

func StartPrint(url string, filePath string) {

	file, _ := os.Open(filePath)
	defer file.Close()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	part, _ := writer.CreateFormFile("file", filepath.Base(file.Name()))
	io.Copy(part, file)

	writer.WriteField("path", "/")
	writer.WriteField("print", "true")
	writer.Close()

	log.Println(fmt.Sprintf("sending post request %s", url))

	r, _ := http.NewRequest("POST", url, body)
	r.Header.Add("Content-Type", writer.FormDataContentType())
	r.Header.Add("X-Api-Key", "F84833BA940E49648F606EA561E97783")

	client := &http.Client{}
	client.Do(r)
}
