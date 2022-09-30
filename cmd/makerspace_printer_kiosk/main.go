package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/configuration"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/middlewares"
	"github.com/maker-space-experimenta/printer-kiosk/internal/common/tasks"

	"github.com/maker-space-experimenta/printer-kiosk/internal/files"
	"github.com/maker-space-experimenta/printer-kiosk/internal/printers"
)

func notFound(w http.ResponseWriter, r *http.Request) {

	log.Fatalf("not found: %v ", r.URL)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text")
	w.Write([]byte("not found"))
}

func main() {

	config, err := configuration.LoadConfig("./config.yml")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	printerRepository := printers.NewPrinterRepository(*config)
	filesRepository := files.NewFileRepository()

	log.Println("config loaded")
	log.Print(config)
	log.Println("")

	router := mux.NewRouter()

	// files_handlers.NewFilesHandler

	printersHandler := printers.NewPrintersHandler(*config)
	// printHandler := internal.printers routes.NewPrintHandler(*config)
	// printersHandler := routes.NewPrintersHandler(*config)
	// filesHandler := routes.NewFilesHandler(*config)
	// octoMockHandler := routes.NewOctoMockHandler(*config)
	// spaHandler := routes.NewSpaHandler(*config, "static", "index.html")

	// router.Path("/api/version").Methods("GET").HandlerFunc(octoMockHandler.GetVersionOctoMock)

	// router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	// router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)

	router.Path("/api/printers").Methods("GET").HandlerFunc(printersHandler.GetPrinters)

	files.AddRoutes(router)

	// router.Path("/api/print").Methods("POST").HandlerFunc(printHandler.PostPrint)

	router.Path("/metrics").Handler(promhttp.Handler())

	// router.PathPrefix("/").Handler(spaHandler)

	router.NotFoundHandler = http.HandlerFunc(notFound)

	// http.Handle("/", router)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{})
	// n.Use(&toolbox_middleware.AuthMiddleware{})
	n.UseHandler(router)

	log.Printf("starting server on port %v", config.Server.Port)

	taskRunner := tasks.NewTaskRunner(*&config.Tasks.Duration)
	taskRunner.AddTask(printerRepository.UpdatePrinters)
	taskRunner.AddTask(filesRepository.UpdateFiles)

	taskRunner.Start()

	//start and listen to requests
	http.ListenAndServe(fmt.Sprintf(":%v", config.Server.Port), n)
}
