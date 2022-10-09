package logs

import (
	"net/http"
	"time"

	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

type chilogger struct {
	logger *zap.SugaredLogger
	name   string
}

// NewHTTPMiddleware returns a new HTTP Middleware handler.
func NewHTTPMiddleware(logger *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return chilogger{
		logger: logger,
		name:   "httpServer",
	}.middleware
}

func (c chilogger) middleware(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		var requestID string

		if reqID := r.Context().Value(middleware.RequestIDKey); reqID != nil {
			v, ok := reqID.(string)
			if ok {
				requestID = v
			}
		}

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)

		next.ServeHTTP(ww, r)

		latency := time.Since(start)

		if c.logger != nil {
			fields := []interface{}{
				"status", ww.Status(),
				"took", latency,
				"request", r.RequestURI,
				"method", r.Method,
			}

			if requestID != "" {
				fields = append(fields, "request-id", requestID)
			}

			c.logger.Infow("request completed", fields...)
		}
	}

	return http.HandlerFunc(fn)
}
