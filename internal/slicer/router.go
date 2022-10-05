package slicer

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

	printersHandler := NewSlicerHandler(*config)

	router.Path("/api/slicer/jobs").Methods("GET").HandlerFunc(printersHandler.GetJobs)
	router.Path("/api/slicer").Methods("POST").HandlerFunc(printersHandler.PostSlicejob)
}
