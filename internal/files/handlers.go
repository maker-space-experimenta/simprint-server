package files

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

type FilesApiResponse struct {
	Data  []PrusaSlicerGcodeMetaData `json:"data"`
	Count int                        `json:"count"`
}

type FilesHandler struct {
	config   configuration.Config
	fileRepo *FileRepository
}

func NewFilesHandler(config configuration.Config) *FilesHandler {
	return &FilesHandler{
		config:   config,
		fileRepo: NewFileRepository(),
	}
}

func (m *FilesHandler) GetFiles(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter files route endpoint GetFiles for " + string(r.URL.Path))

	filesList := m.fileRepo.Files

	if len(filesList) == 0 {
		filesList = make([]PrusaSlicerGcodeMetaData, 0)
	}

	response := FilesApiResponse{
		Data:  filesList,
		Count: len(filesList),
	}

	jsonResp, err := json.Marshal(response)
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

func (m *FilesHandler) PostFiles(w http.ResponseWriter, r *http.Request) {
	log.Printf("enter files route endpoint PostFile for " + string(r.URL.Path))

	file, handler, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dirName := m.config.Files.TempDir

	err = os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		panic(err)
	}

	url := fmt.Sprintf("%v/%v", dirName, handler.Filename)
	log.Printf(url)

	f, err := os.OpenFile(url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	_, _ = io.WriteString(w, "File "+url+" Uploaded successfully")
	_, _ = io.Copy(f, file)
}
