package application

import (
	"context"
	"fmt"
	"github.com/aivanov/game/http/server"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
)

func Run(ctx context.Context) int {
	// Creating a logger with settings for production
	logger := setupLogger()

	shutDownFunc, err := server.Run(ctx, logger, 3, 3)
	if err != nil {
		logger.Error(err.Error())

		return 1 // return the code for registering an error with the system
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	<-c
	cancel()
	// shut down the server
	shutDownFunc(ctx)

	return 0

}

// logger settings
func setupLogger() *zap.Logger {
	// Configuring the logger
	config := zap.NewProductionConfig()

	// Logging level
	config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)

	// Setting up a logger with config
	logger, err := config.Build()
	if err != nil {
		fmt.Printf("Logger setup error: %v\n", err)
	}

	return logger
}
