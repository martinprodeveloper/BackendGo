// Package main
package main

import (
	"log/slog"
	"net/http"
	"os"
	"time"

	_ "github.com/martinprodeveloper/devboard/docs"
	"github.com/martinprodeveloper/devboard/internal/handler"
	"github.com/martinprodeveloper/devboard/internal/logger"
	"github.com/martinprodeveloper/devboard/internal/middlewares"
	"github.com/martinprodeveloper/devboard/internal/server"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Devboard API
// @version 1.0
// @descrption API REST para gestión de proyectos
// @contact.name @martinprodeveloper
// @host localhost:8080
// @BasePath /api/v1
func main() {

	// 1. Entorno
	env := os.Getenv("GO_ENV")
	slog.Info("entorno cargado", "GO_ENV", env)

	// 2. Logger
	var log *slog.Logger
	if env == "production" {
		log = logger.New(logger.ProductionConfig())
	} else {
		log = logger.New(logger.DefaultConfig())
	}

	// 3. Dependencies compartidas
	// validate := validator.New()
	// notifier := notification.NewLogNotifier(log)
	// _ = notifier

	// 4. Servidor con functional options
	srv := server.New(":8080",
		server.WithLogger(log),
		server.WithReadTimeout(15*time.Second),
		server.WithWriteTimeout(30*time.Second),
	)

	// 5. Handlers
	healthHandler := handler.NewHealthHandler()

	// 6. Registro de rutas
	srv.RegisterRoutes("GET /docs/", httpSwagger.WrapHandler)
	srv.RegisterRoutes("GET /health", healthHandler)
	srv.RegisterRoutes("GET /panic", http.HandlerFunc(func(_ http.ResponseWriter, _ *http.Request) {
		panic("Simulated panic for testing")
	}))

	// 7. Middlewares
	srv.UseChain(
		middlewares.Recovery(log),
		middlewares.Recovery(log),
	)

	// 8. Arrancar shutdown limpio

	if err := srv.Start(); err != nil {
		log.Error("Error fatal", "error", err)
		os.Exit(1)
	}
}
