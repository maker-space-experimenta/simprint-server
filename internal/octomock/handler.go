package octomock

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

type OctoMockResponse struct {
}

type OctoMockHandler struct {
	config configuration.Config
}

func NewOctoMockHandler(config configuration.Config) *OctoMockHandler {
	return &OctoMockHandler{
		config: config,
	}
}

type OctoMockVersionResponse struct {
	Api    string `json:"api"`
	Server string `json:"server"`
	Text   string `json:"text"`
}

func (m *OctoMockHandler) GetVersionOctoMock(w http.ResponseWriter, r *http.Request) {

	result := OctoMockVersionResponse{
		Api:    "0.1",
		Server: "1.3.10",
		Text:   "OctoPrint 1.3.10",
	}

	jsonResp, err := json.Marshal(result)
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
