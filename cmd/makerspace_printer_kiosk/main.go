package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/middlewares"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/tasks"
	"github.com/maker-space-experimenta/printer-kiosk/internal/jobs"
	"github.com/maker-space-experimenta/printer-kiosk/internal/octomock"
	"github.com/maker-space-experimenta/printer-kiosk/internal/slicer"
	"github.com/maker-space-experimenta/printer-kiosk/internal/spa"

	"github.com/maker-space-experimenta/printer-kiosk/internal/files"
	"github.com/maker-space-experimenta/printer-kiosk/internal/printers"
)

func notFound(w http.ResponseWriter, r *http.Request) {
	logger := logging.NewLogger()
	logger.Infof("not found: %v ", r.URL)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text")
	w.Write([]byte("not found"))
}

func main() {

	configService := configuration.NewConfigService()
	err := configService.LoadConfig("./config.yml")
	if err != nil {
		log.Fatalf("cannot load config:", err)
	}

	config, err := configService.GetConfig()
	if err != nil {
		log.Fatalf("cannot load config:", err)
	}

	logger := logging.NewLogger()
	logger.Infof("starting application printerkiosk-api")

	router := mux.NewRouter()
	octomock.AddRoutes(router)
	files.AddRoutes(router)
	printers.AddRoutes(router)
	jobs.AddRoutes(router)
	slicer.AddRoutes(router)
	spa.AddRoutes(router)
	router.Path("/metrics").Handler(promhttp.Handler())
	router.NotFoundHandler = http.HandlerFunc(notFound)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{})
	// n.Use(&toolbox_middleware.AuthMiddleware{})
	n.UseHandler(router)

	logger.Infof("starting server on port %v", config.Server.Port)

	printerRepository := printers.NewPrinterRepository(*config)
	filesRepository := files.NewFileRepository(*config)
	taskRunner := tasks.NewTaskRunner(*&config.Tasks.Duration)
	taskRunner.AddTask(printerRepository.UpdatePrinters)
	taskRunner.AddTask(filesRepository.UpdateFiles)

	taskRunner.Start()

	//start and listen to requests
	http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), n)
}
