package simplehttpcounter

import (
	"fmt"
	"net/http"
)

func HTTPCounter(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HTTP Request Counter")
		next.ServeHTTP(w, r)
	})
}
