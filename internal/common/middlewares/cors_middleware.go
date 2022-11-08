package middlewares

import (
	"net/http"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type CorsMiddleware struct{}

func (*CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := logging.NewLogger()

	logger.Infof("The Cors middleware is executing!")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	next.ServeHTTP(w, r)
}
