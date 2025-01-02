package middleware

import (
	"net/http"

	"go.opentelemetry.io/otel/attribute"
	"proj/pkg/tracer"
)

func Trace() func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			operationName := "http_" + r.Method + "_" + r.URL.Path
			ctx, span := tracer.StartSpan(r.Context(), operationName)
			defer span.End()

			span.SetAttributes(
				attribute.String("http.method", r.Method),
				attribute.String("http.url", r.URL.String()),
				attribute.String("http.user_agent", r.UserAgent()),
				attribute.String("http.remote_addr", r.RemoteAddr),
			)

			rw := &responseWriter{ResponseWriter: w}

			next(rw, r.WithContext(ctx))

			span.SetAttributes(
				attribute.Int("http.status_code", rw.statusCode),
			)
		}
	}
}

type responseWriter struct {
	http.ResponseWriter
	statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
	rw.statusCode = code
	rw.ResponseWriter.WriteHeader(code)
}
