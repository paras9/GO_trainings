package middleware

import (
	"net/http"
	"time"

	"csv_manager/utils" // Import the utils package
)

// LoggingMiddleware logs request details using the utils.Logger
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		utils.Logger.Printf("Started %s %s", r.Method, r.URL.Path) // Use utils.Logger
		next.ServeHTTP(w, r)
		utils.Logger.Printf("Completed %s in %v", r.URL.Path, time.Since(start))
	})
}
