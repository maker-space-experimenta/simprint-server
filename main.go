package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/maker-space-experiemnta/printer-kiosk/middlewares"
	"github.com/maker-space-experiemnta/printer-kiosk/routes"
	"github.com/maker-space-experiemnta/printer-kiosk/util"
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

	log.Println("config loaded")
	log.Print(config)
	log.Println("")

	router := mux.NewRouter()

	// ordersRepo := toolbox_repositories.OrderRepo{}

	filesHandler := routes.NewFilesHandler(config)
	octoMockHandler := routes.NewOctoMockHandler(config)

	// router.Path("/metrics").Methods("GET").HandlerFunc(toolbox_routes.GetMetrics)

	router.Path("/api/version").Methods("GET").HandlerFunc(octoMockHandler.GetVersionOctoMock)

	router.Path("/api/files/{location}").Methods("GET").HandlerFunc(filesHandler.GetFiles)
	router.Path("/api/files/{location}").Methods("POST").HandlerFunc(filesHandler.PostFiles)

	router.Path("/api/readfile/{filename}").Methods("GET").HandlerFunc(filesHandler.ReadFile)

	router.Path("/metrics").Handler(promhttp.Handler())

	router.NotFoundHandler = http.HandlerFunc(notFound)

	// http.Handle("/", router)

	n := negroni.New()
	n.Use(&middlewares.CorsMiddleware{})
	n.Use(&middlewares.LoggerMiddleware{})
	// n.Use(&toolbox_middleware.AuthMiddleware{})
	n.UseHandler(router)

	log.Printf("starting server on port %v", config.Port)

	//start and listen to requests
	http.ListenAndServe(fmt.Sprintf(":%v", config.Port), n)
}
