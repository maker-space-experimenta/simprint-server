package spa

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

	spaHandler := NewSpaHandler(*config, "static", "index.html")
	router.PathPrefix("/").Handler(spaHandler)
}
