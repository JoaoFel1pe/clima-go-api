package middlewares

import (
	"fmt"
	"net/http"
)

type statusRecorder struct {
	http.ResponseWriter
	status int
}

func (rec *statusRecorder) WriteHeader(code int) {
	rec.status = code
	rec.ResponseWriter.WriteHeader(code)
}

func LogginMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rec := statusRecorder{ResponseWriter: w, status: http.StatusOK}
		next.ServeHTTP(&rec, r)
		statusText := http.StatusText(rec.status)
		color := "\033[32m" // Green
		if rec.status >= 400 {
			color = "\033[31m" // Red
		}
		reset := "\033[0m"
		fmt.Printf("%sRequest: %s %s %d %s%s\n", color, r.Method, r.RequestURI, rec.status, statusText, reset)
	})
}
