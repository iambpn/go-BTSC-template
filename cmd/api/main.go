package main

import (
	"context"
	"fmt"
	"github.com/iambpn/go-http-template/internal/logger"
	"github.com/iambpn/go-http-template/internal/server"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	_ "github.com/joho/godotenv/autoload"
)

func getLogFileName() string {
	fileName := os.Getenv("LOG_FILE")

	if fileName == "" {
		fileName = "application.log"
	}

	return fileName
}

func openLogFile(filename string) *os.File {
	logFile, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		slog.Error("error opening log file.", "error", err)
		os.Exit(1)
	}

	return logFile
}

func main() {
	ctx := context.Background()
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	// setup logger
	logFile := openLogFile(getLogFileName())
	defer logFile.Close()
	logger.SetupLogger(logFile)

	// Initiate server
	server := server.NewServer()

	// start server on different go-routine
	go func() {
		fmt.Printf("starting server on addr: %s", server.Addr)

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Log.Error(fmt.Sprintf("error while starting server: %s\n", err))
			os.Exit(1)
			return
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		// wait until interrupt signal triggers
		<-ctx.Done()

		// Make a new context for the Shutdown
		shutdownCtx := context.Background()

		// Terminate context after some timeout
		shutdownCtx, cancel := context.WithTimeout(shutdownCtx, 10*time.Second)

		defer cancel()

		if err := server.Shutdown(shutdownCtx); err != nil {
			fmt.Printf("error shutting down http server: %s\n", err)
		}
	}()

	wg.Wait()
}
