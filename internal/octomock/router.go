package octomock

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {

	logger := logging.NewLogger()
	logger.Debugf("Running OctoMock AddRoutes")

	octoMockHandler := NewOctoMockHandler(config)
	logger.Debugf("OctoMock Handler created")

	router.Path("/api/version").Methods("GET").HandlerFunc(octoMockHandler.GetVersionOctoMock)

	// router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	// router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)
}
