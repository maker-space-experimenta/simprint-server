package octomock

import (
	"encoding/json"
	"net/http"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type OctoMockResponse struct {
}

type OctoMockHandler struct {
	config *configuration.Config
	logger *logging.Logger
}

func NewOctoMockHandler(config *configuration.Config) *OctoMockHandler {

	logger := logging.NewLogger()
	logger.Debugf("Create new Mock Handler")

	if config == nil {
		logger.Errorf("no config found")
	}

	return &OctoMockHandler{
		config: config,
		logger: logger,
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
		m.logger.Errorf("Error happened in JSOn marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
