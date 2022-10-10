package files

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
		logger.Errorf("cannot load config:", err)
	}

	filesHandler := NewFilesHandler(*config)

	router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)
}
