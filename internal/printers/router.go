package printers

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {
	logger := logging.NewLogger()
	logger.Debugf("Running printers AddRoutes")

	printersHandler := NewPrintersHandler(*config)

	router.Path("/api/printers").Methods("GET").HandlerFunc(printersHandler.GetPrinters)
}
