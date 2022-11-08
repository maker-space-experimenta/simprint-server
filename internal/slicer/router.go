package slicer

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {
	logger := logging.NewLogger()
	logger.Debugf("Running slicer AddRoutes")

	printersHandler := NewSlicerHandler(*config)

	router.Path("/api/slicer/jobs").Methods("GET").HandlerFunc(printersHandler.GetJobs)
	router.Path("/api/slicer").Methods("POST").HandlerFunc(printersHandler.PostSlicejob)
}
