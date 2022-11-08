package spa

import (
	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

func AddRoutes(router *mux.Router, config *configuration.Config) {
	logger := logging.NewLogger()
	logger.Debugf("Running SPA AddRoutes")

	spaHandler := NewSpaHandler(*config, "static", "index.html")
	router.PathPrefix("/").Handler(spaHandler)
}
