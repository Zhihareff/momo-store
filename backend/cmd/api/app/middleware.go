package app

import (
	"net/http"
	"time"
)

// TimingsMiddleware records timings of app handlers
func (i *Instance) TimingsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		i.responseTimings.
			With(map[string]string{"handler": r.URL.Path}).
			Observe(float64(time.Since(start).Milliseconds()))
	})
}

// RequestsMiddleware records count requests to app handlers
func (i *Instance) RequestsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		i.requestCounter.Inc()
		next.ServeHTTP(w, r)
	})
}
