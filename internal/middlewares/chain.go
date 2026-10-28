// Package middlewares provides chain.
package middlewares

import (
	"net/http"
)

// Middelware func
type Middelware func(http.Handler) http.Handler

// Chain func
func Chain(handler http.Handler, middlewares ...Middelware) http.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](handler)
	}
	return handler
}
