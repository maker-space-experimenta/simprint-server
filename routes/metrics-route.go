package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type MetricResponse struct {
	Printers int
}

type MetricsHandler struct {
}

func NewMetricsHandler() *MetricsHandler {
	return &MetricsHandler{}
}

func (m *MetricsHandler) GetMetrics(w http.ResponseWriter, r *http.Request) {

	resp := MetricResponse{
		Printers: 2,
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSOn marshal. Err %s", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(nil)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResp)
}
