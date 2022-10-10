package printers

import (
	"encoding/json"
	"net/http"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type PrintersHandler struct {
	config      configuration.Config
	printerRepo *PrinterRepository
	logger      *logging.Logger
}

func NewPrintersHandler(config configuration.Config) *PrintersHandler {
	return &PrintersHandler{
		config:      config,
		printerRepo: NewPrinterRepository(config),
		logger:      logging.NewLogger(),
	}
}

func (m *PrintersHandler) GetPrinters(w http.ResponseWriter, r *http.Request) {
	m.logger.Infof("running GetPrinters")

	jsonResp, err := json.Marshal(m.printerRepo.Printers)
	if err != nil {
		m.logger.Errorf("Error happened in JSON marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
