package middlewares

import (
	"net/http"
)

// We know we will only be replying in json, lets clean up redudant code.
func JSONResponse(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
