// internal/server/server.go
package server

import (
	"net/http"

	"tic2/internal/config"
	"tic2/internal/middleware"
	"tic2/internal/web/handler"
)

func New(
	cfg config.Config,
	h *handler.GameHandler,
	authHandler *handler.AuthHandler,
	authenticator *middleware.UserAuthenticator,
	protectedMux := http.NewServeMux()
	h.RegisterRoutes(protectedMux)
	userHandler.RegisterRoutes(protectedMux) // новый
	protectedHandler := authenticator.Authenticate(protectedMux)
) *http.Server {
	// публичные маршруты
	publicMux := http.NewServeMux()
	publicMux.HandleFunc("POST /signup", authHandler.SignUp)
	publicMux.HandleFunc("POST /signin", authHandler.SignIn)

	// защищённые маршруты
	protectedMux := http.NewServeMux()
	h.RegisterRoutes(protectedMux)
	protectedHandler := authenticator.Authenticate(protectedMux)

	// корневой mux
	rootMux := http.NewServeMux()
	rootMux.Handle("/signup", publicMux)
	rootMux.Handle("/signin", publicMux)
	rootMux.Handle("/", protectedHandler)

	var root http.Handler = rootMux
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
