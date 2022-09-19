package middlewares

import (
	"fmt"
	"net/http"
	"time"
)

type LoggerMiddleware struct{}

func (*LoggerMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("The logger middleware is executing!")
	t := time.Now()
	next.ServeHTTP(w, r)

	// for k, v := range r.Header {
	// 	fmt.Printf(" %q, Value %q\n", k, v)
	// }

	fmt.Printf("Execution time: %s \n", time.Now().Sub(t).String())
}
