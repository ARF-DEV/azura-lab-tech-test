package middleware

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tech-test-azura-lab/services"
)

func Method(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			w.WriteHeader(http.StatusMethodNotAllowed)
			w.Header().Set("Content-Type", "application/json")
			resp, _ := json.Marshal(services.ErrorResponse{Err: fmt.Sprintf("Only method %s is allowed", method)})
			w.Write(resp)
			return
		}
		next.ServeHTTP(w, r)
	})
}
func ContentType(value string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", value)
		next.ServeHTTP(w, r)
	})
}
