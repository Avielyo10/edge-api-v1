package main

import (
	"context"
	"github.com/Avielyo10/edge-api/internal/common/logs"
	"github.com/Avielyo10/edge-api/internal/common/server"
	"github.com/Avielyo10/edge-api/internal/edge/ports/images"
	"github.com/Avielyo10/edge-api/internal/edge/service"
	"github.com/go-chi/chi/v5"
	"github.com/redhatinsights/edge-api/config"
	"net/http"
)

func main() {
	config.Init() // init config
	logs.Init()   // init logger

	ctx := context.Background()

	application := service.NewApplication(ctx)

	server.RunHTTPServer(config.Get(), func(router chi.Router) http.Handler {
		return ports.HandlerFromMux(
			ports.NewHttpServer(application),
			router,
		)
	})
}
