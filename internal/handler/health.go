// Package handler provides HTTP handlers for the application.
package handler

import (
	"encoding/json"
	"net/http"
)

// HealthHandler struct
type HealthHandler struct{}

// NewHealthHandler constructor
func NewHealthHandler() *HealthHandler {
	return &HealthHandler{}
}

// HealthResponse struct
type HealthResponse struct {
	Status  string `json:"status"`
	Version string `json:"version"`
}

// ServeHTTP func
// @Summary Health Check
// @Description Verifica que el servidor está corriendo
// @Tags system
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (handler *HealthHandler) ServeHTTP(writer http.ResponseWriter, _ *http.Request) {
	response := HealthResponse{
		Status:  "ok",
		Version: "1.0.0",
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(writer).Encode(response); err != nil {
		return
	}
}
