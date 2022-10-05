package printers

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

func AddRoutes(router *mux.Router) {

	configService := configuration.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		log.Printf("FATAL: cannot load config: %v", err)
	}

	printersHandler := NewPrintersHandler(*config)

	router.Path("/api/printers").Methods("GET").HandlerFunc(printersHandler.GetPrinters)
}
