package server

import (
	"context"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"github.com/noodlensk/task-tracker/internal/common/auth"
	"github.com/noodlensk/task-tracker/internal/common/logs"
)

func RunHTTPServer(ctx context.Context, addr string, logger *zap.SugaredLogger, createHandler func(router chi.Router) http.Handler) error {
	return RunHTTPServerOnAddr(ctx, addr, logger, createHandler)
}

func RunHTTPServerOnAddr(ctx context.Context, addr string, logger *zap.SugaredLogger, createHandler func(router chi.Router) http.Handler) error {
	apiRouter := chi.NewRouter()
	setAPIMiddlewares(apiRouter, logger)

	rootRouter := chi.NewRouter()
	// we are mounting all APIs under /api path
	rootRouter.Mount("/api", createHandler(apiRouter))

	logger.Infof("Starting HTTP server on %q", addr)

	errCh := make(chan error)

	srv := &http.Server{Addr: addr, Handler: rootRouter} //nolint:gosec

	go func() { errCh <- srv.ListenAndServe() }()

	select {
	case err := <-errCh:
		return err
	case <-ctx.Done():
		logger.Info("Stopping HTTP server")

		newCtx, cancel := context.WithTimeout(ctx, time.Second*5)
		defer cancel()

		return srv.Shutdown(newCtx)
	}
}

func setAPIMiddlewares(router *chi.Mux, logger *zap.SugaredLogger) {
	router.Use(middleware.RequestID)
	router.Use(middleware.RealIP)
	router.Use(middleware.Recoverer)
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	router.Use(
		middleware.SetHeader("X-Content-Type-Options", "nosniff"),
		middleware.SetHeader("X-Frame-Options", "deny"),
		middleware.SetHeader("Content-Type", "application/json"),

		middleware.NoCache,
	)

	router.Use(logs.NewHTTPMiddleware(logger))
	router.Use(auth.JWTHttpMiddleware{Secret: "secret", AuthURL: "/api/auth/login"}.Middleware)
}
