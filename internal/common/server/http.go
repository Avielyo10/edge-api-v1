package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/redhatinsights/edge-api/config"
	"github.com/redhatinsights/platform-go-middlewares/identity"
	"github.com/redhatinsights/platform-go-middlewares/request_id"
	"github.com/sirupsen/logrus"
)

// RunHTTPServer runs an http server.
func RunHTTPServer(cfg *config.EdgeConfig, createHandler func(router chi.Router) http.Handler) {
	apiRouter := chi.NewRouter()
	setMiddlewares(cfg, apiRouter)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api/edge/v1 path
	rootRouter.Mount("/api/edge/v1", createHandler(apiRouter))

	logrus.Info("Starting HTTP server")

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.WebPort),
		Handler:      rootRouter,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		logs.LogErrorAndPanic("web service stopped unexpectedly", err)
	}
}

func setMiddlewares(cfg *config.EdgeConfig, router *chi.Mux) {
	if cfg.Auth {
		router.Use(identity.EnforceIdentity)
	}
	router.Use(
		request_id.ConfiguredRequestID("x-rh-insights-request-id"),
		middleware.RequestID,
		middleware.RealIP,
		middleware.Recoverer,
		middleware.NoCache,
		logs.NewStructuredLogger(logrus.StandardLogger()),
	)
	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
	)
}
