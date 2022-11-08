package middlewares

import (
	"net/http"
	"time"

	"github.com/maker-space-experimenta/printer-kiosk/internal/common/logging"
)

type LoggerMiddleware struct{}

func (*LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	logger := logging.NewLogger()

	logger.Debugf("The logger middleware is executing!")
	t := time.Now()
	next.ServeHTTP(w, r)

	// for k, v := range r.Header {
	// 	fmt.Printf(" %q, Value %q", k, v)
	// }

	logger.Debugf("Execution time: %s ", time.Now().Sub(t).String())
}
