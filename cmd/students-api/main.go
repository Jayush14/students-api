package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Jayush14/students-api/internal/config"
	"github.com/Jayush14/students-api/internal/http/handlers/student"
)

func main() {
	//load config
	cfg := config.MustLoad()
	//database setup
	//router setup
	router := http.NewServeMux()

	router.HandleFunc("POST /api/students", student.New())
	//http server setup
	server := http.Server{
		Addr:    cfg.HttpServer.Address,
		Handler: router,
	}
	slog.Info("Starting server on " + cfg.HttpServer.Address)

	done := make(chan os.Signal, 1)

	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			slog.Error("Failed to start server", "error", err)
		}
	}()

	<-done

	slog.Info("Stopping server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := server.Shutdown(ctx)
	if err != nil {
		slog.Error("Failed to stop server", "error", err)
	}

	slog.Info("Server stopped")

}
