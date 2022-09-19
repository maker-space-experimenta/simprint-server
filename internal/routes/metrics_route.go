package routes

import (
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

func (m *MetricsHandler) GetCycleSpeed(w http.ResponseWriter, r *http.Request) {
	b := []byte("{ content: \"test\" }")

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	w.Write(b)
}
