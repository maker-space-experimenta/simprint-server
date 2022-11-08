package jobs

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {
	logger := logging.NewLogger()
	logger.Debugf("Running jobs AddRoutes")

	jobHandler := NewPrintHandler(*config)
	router.PathPrefix("/api/jobs").Methods("POST").HandlerFunc(jobHandler.PostPrint)
}
