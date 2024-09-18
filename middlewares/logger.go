package middlewares

import (
	"log"
	"net/http"
	"time"
)

// Simple logger middleware to show what calls are being handled
func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()
		log.Printf("Started %s %s", r.Method, r.URL.Path)

		rr := &responseRecorder{w, http.StatusOK}
		next.ServeHTTP(rr, r)

		log.Printf("Completed %s %s in %v with status %d", r.Method, r.URL.Path, time.Since(startTime), rr.statusCode)
	})
}

type responseRecorder struct {
	http.ResponseWriter
	statusCode int
}

func (rr *responseRecorder) WriteHeader(code int) {
	rr.statusCode = code
	rr.ResponseWriter.WriteHeader(code)
}
