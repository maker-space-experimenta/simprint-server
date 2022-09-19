package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maker-space-experimenta/printer-kiosk/internal/middlewares"
	"github.com/maker-space-experimenta/printer-kiosk/internal/repositories"
	"github.com/maker-space-experimenta/printer-kiosk/internal/routes"
	"github.com/maker-space-experimenta/printer-kiosk/internal/util"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/negroni"
)

func notFound(w http.ResponseWriter, r *http.Request) {

	log.Fatalf("not found: %v ", r.URL)

	w.WriteHeader(http.StatusNotFound)
	w.Header().Set("Content-Type", "text")
	w.Write([]byte("not found"))
}

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	printerRepository := repositories.NewPrinterRepository(config)
	filesRepository := repositories.NewFileRepository(config)

	log.Println("config loaded")
	log.Print(config)
	log.Println("")

	router := mux.NewRouter()

	// ordersRepo := toolbox_repositories.OrderRepo{}

	metricsHandler := routes.NewMetricsHandler()
	printHandler := routes.NewPrintHandler(config)
	printersHandler := routes.NewPrintersHandler(config)
	filesHandler := routes.NewFilesHandler(config, *filesRepository)
	octoMockHandler := routes.NewOctoMockHandler(config)
	spaHandler := routes.NewSpaHandler(config, "static", "index.html")

	router.Path("/api/version").Methods("GET").HandlerFunc(octoMockHandler.GetVersionOctoMock)

	router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)

	router.Path("/api/printers").Methods("GET").HandlerFunc(printersHandler.GetPrinters)

	router.Path("/api/print").Methods("POST").HandlerFunc(printHandler.PostPrint)

	router.Path("/metrics").Handler(promhttp.Handler())
	router.Path("/api/metrics/speed").HandlerFunc(metricsHandler.GetCycleSpeed)

	router.PathPrefix("/").Handler(spaHandler)

	router.NotFoundHandler = http.HandlerFunc(notFound)

	// http.Handle("/", router)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{})
	// n.Use(&toolbox_middleware.AuthMiddleware{})
	n.UseHandler(router)

	log.Printf("starting server on port %v", config.Port)

	taskRunner := util.NewTaskRunner(config)
	taskRunner.AddTask(printerRepository.UpdatePrinters)
	taskRunner.AddTask(filesRepository.UpdateFiles)

	taskRunner.Start()

	//start and listen to requests
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), n)
}
