// Package middlewares provides middleware functions for HTTP servers.
package middlewares

import (
	"log/slog"
	"net/http"
	"runtime/debug"
)

// Recovery middleware recovers from panics and logs the error and stack trace.
func Recovery(logger *slog.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					stack := debug.Stack()
					logger.ErrorContext(request.Context(), "panic recover",
						slog.Any("error", err),
						slog.String("stack", string(stack)),
						slog.String("path", request.URL.Path),
					)
					http.Error(writer, "Error interno del servidor", http.StatusInternalServerError)
				}
			}()

			next.ServeHTTP(writer, request)
		})
	}
}
