package files

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"path"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/helper"
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
		fileRepo: NewFileRepository(config),
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

	filepath, err := helper.SaveFileFromForm(r, "file", path.Join(m.config.Files.TempDir, "gcode"), "sliced.gcode")
	if err != nil {
		log.Printf("FATAL: could not load and save file, %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	_, _ = io.WriteString(w, "File "+filepath+" Uploaded successfully")
}
