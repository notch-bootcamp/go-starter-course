package middleware

import (
	"fmt"
	"net/http"
)

func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	}
}
