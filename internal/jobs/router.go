package jobs

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

	jobHandler := NewPrintHandler(*config)
	router.PathPrefix("/api/jobs").Methods("POST").HandlerFunc(jobHandler.PostPrint)
}
