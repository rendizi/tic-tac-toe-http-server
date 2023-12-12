package server

import (
	"context"
	"fmt"
	"github.com/aivanov/game/http/server/handler"
	"github.com/aivanov/game/internal/service"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// routing
func new(ctx context.Context,
	logger *zap.Logger,
	lifeService service.LifeService,
) (http.Handler, error) {
	muxHandler, err := handler.New(ctx, lifeService)
	if err != nil {
		return nil, fmt.Errorf("handler initialization error: %w", err)
	}
	// middleware for handlers
	muxHandler = handler.Decorate(muxHandler, loggingMiddleware(logger))

	return muxHandler, nil
}

func Run(
	ctx context.Context,
	logger *zap.Logger,
	height, width int,
) (func(context.Context) error, error) {
	// service with game
	lifeService, err := service.New()
	if err != nil {
		return nil, err
	}

	muxHandler, err := new(ctx, logger, *lifeService)
	if err != nil {
		return nil, err
	}

	srv := &http.Server{Addr: ":8081", Handler: muxHandler}

	go func() {
		// Start the server
		if err := srv.ListenAndServe(); err != nil {
			logger.Error("ListenAndServe",
				zap.String("err", err.Error()))
		}
	}()
	// return the function to shut down the server
	return srv.Shutdown, nil
}

// middleware for logging requests
func loggingMiddleware(logger *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			start := time.Now()

			// Skip the request to the next handler
			next.ServeHTTP(w, r)

			// Finish logging after request execution
			duration := time.Since(start)
			logger.Info("HTTP request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Duration("duration", duration),
			)
		})
	}
}
