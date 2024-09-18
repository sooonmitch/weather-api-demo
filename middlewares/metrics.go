package middlewares

import (
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	mu           sync.Mutex
	requestCount int
	totalTime    time.Duration
)

// Demo of metrics middleware, could throw this out to Prometheus or our own metrics endpoint but currently just log it for simplicity
func Metrics(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		elapsed := time.Since(startTime)
		mu.Lock()
		requestCount++
		totalTime += elapsed
		averageTime := totalTime / time.Duration(requestCount)
		mu.Unlock()

		log.Printf("Request count: %d, Average response time: %v", requestCount, averageTime)
	})
}
