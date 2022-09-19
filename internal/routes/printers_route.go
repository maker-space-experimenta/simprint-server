package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/maker-space-experimenta/printer-kiosk/internal/repositories"
	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
)

type PrintersHandler struct {
	config      util.Config
	printerRepo repositories.PrinterRepository
	// printers []models.Printer
}

func NewPrintersHandler(config util.Config) *PrintersHandler {
	return &PrintersHandler{
		config: config,
	}
}

func (m *PrintersHandler) GetPrinters(w http.ResponseWriter, r *http.Request) {
	log.Printf("running GetPrinters")

	jsonResp, err := json.Marshal(m.printerRepo.Printers)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
