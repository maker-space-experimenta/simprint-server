package printers

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router) {
	logger := logging.NewLogger()

	configService := configuration.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		logger.Errorf("FATAL: cannot load config: %v", err)
	}

	printersHandler := NewPrintersHandler(*config)

	router.Path("/api/printers").Methods("GET").HandlerFunc(printersHandler.GetPrinters)
}
