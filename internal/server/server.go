package server

import (
	"net/http"

	"tic2/internal/config"
	"tic2/internal/middleware"
	"tic2/internal/web/handler"
)

func New(cfg config.Config, h *handler.GameHandler) *http.Server {
	mux := http.NewServeMux()
	h.RegisterRoutes(mux)
	var root http.Handler = mux
	root = middleware.Logging(root)
	root = middleware.Recovery(root)

	return &http.Server{
		Addr:         cfg.HTTPAddr,
		Handler:      root,
		ReadTimeout:  cfg.ReadTimeout,
		WriteTimeout: cfg.WriteTimeout,
		IdleTimeout:  cfg.IdleTimeout,
	}
}
