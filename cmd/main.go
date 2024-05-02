package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/iambpn/go-http-template/internal/config"
	"github.com/iambpn/go-http-template/internal/server"
)

func main() {
	ctx := context.Background()

	// passing os.Stdout because this will make it easy to test application
	if err := run(ctx, os.Stdout, os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, w io.Writer, args []string) error {
	logger := log.New(w, "", log.Flags())

	config := &config.AppConfig{
		Host: args[1],
		Port: args[2],
	}

	// For handling graceful shutdown,
	// listen to os signals like exit signal
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	srv := server.NewServer(
		logger,
		config,
	)

	httpServer := &http.Server{
		Addr:    net.JoinHostPort(config.Host, config.Port),
		Handler: srv,
	}

	// listen server on different go-routine
	go func() {
		log.Printf("listening on %s\n", httpServer.Addr)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Fprintf(w, "error listening and serving: %s\n", err)
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

		if err := httpServer.Shutdown(shutdownCtx); err != nil {
			fmt.Fprintf(w, "error shutting down http server: %s\n", err)
		}
	}()
	wg.Wait()
	return nil
}
