package middleware

import (
	"github.com/sirupsen/logrus"
	"net/http"
	"oprec/pkg/exception"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(writer http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: writer}
}

func (writer responseWriter) Status() int {
	return writer.status
}

func (writer *responseWriter) WriteHeader(code int) {
	if writer.wroteHeader {
		return
	}

	writer.status = code
	writer.ResponseWriter.WriteHeader(int(code))
	writer.wroteHeader = true

	return
}

func LoggerMiddleware(logger *logrus.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(writer http.ResponseWriter, request *http.Request) {
			start := time.Now()

			// Recover
			defer func() {
				err := recover()
				if err != nil {
					writer.WriteHeader(http.StatusInternalServerError)
					logger.WithFields(logrus.Fields{
						"status":   http.StatusInternalServerError,
						"method":   request.Method,
						"path":     request.URL.EscapedPath(),
						"duration": time.Since(start),
						"error":    err,
					}).Errorln()

					exception.ErrorHandler(writer,request,err)
				}
			}()

			// Logger
			wrapped := wrapResponseWriter(writer)

			next.ServeHTTP(wrapped, request)
			logger.WithFields(logrus.Fields{
				"status":   wrapped.Status(),
				"method":   request.Method,
				"path":     request.URL.EscapedPath(),
				"duration": time.Since(start),
			}).Println()
		}

		return http.HandlerFunc(fn)
	}
}
