package files

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {

	filesHandler := NewFilesHandler(*config)

	router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)
}
