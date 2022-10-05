package octomock

import (
	"log"

	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
)

func AddRoutes(router *mux.Router) {

	configService := configuration.NewConfigService()
	config, err := configService.GetConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	octoMockHandler := NewOctoMockHandler(*config)

	router.Path("/api/version").Methods("GET").HandlerFunc(octoMockHandler.GetVersionOctoMock)

	// router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	// router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)
}
