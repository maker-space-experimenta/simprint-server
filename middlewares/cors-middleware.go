package middlewares

import (
	"fmt"
	"net/http"
)

type CorsMiddleware struct{}

func (*CorsMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	fmt.Println("The Cors middleware is executing!")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	next.ServeHTTP(w, r)
}
