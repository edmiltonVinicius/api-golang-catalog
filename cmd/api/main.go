package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/edmiltonVinicius/go-api-catalog/internal/adapters/httpapi"
	"github.com/edmiltonVinicius/go-api-catalog/internal/adapters/httpapi/handlers"
	"github.com/edmiltonVinicius/go-api-catalog/internal/adapters/postgres"
	"github.com/edmiltonVinicius/go-api-catalog/internal/application/product"
	"github.com/edmiltonVinicius/go-api-catalog/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

func main() {

	// 0. Start logger
	config.StartLogger()
	defer config.Logger.Sync()

	// 1. Load configuration
	cfg, err := config.LoadEnv()
	if err != nil {
		config.Logger.Fatal("failed to load config", zap.Error(err))
	}

	// 2. Root context + OS signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 3. Postgres connection pool
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		config.Logger.Fatal("failed to create pgxpool", zap.Error(err))
	}
	defer pool.Close()

	// 4. Basic health check
	if err := pool.Ping(ctx); err != nil {
		config.Logger.Fatal("failed to ping postgres", zap.Error(err))
	}
	config.Logger.Info("Postgres connection pool created successfully")

	// 5. Adapters (infraestructure)
	postgresRepo := postgres.New(pool)

	// 6. Core (domain services)
	productService := product.NewService(postgresRepo, postgresRepo, postgresRepo)

	// 7. Http handler
	handler := handlers.NewHandler(productService)
	router := httpapi.NewRouter(handler)

	// 8. Http server
	server := &http.Server{
		Addr:         fmt.Sprintf(":%s", cfg.HTTPPort),
		Handler:      router,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// 9. Start server
	go func() {
		config.Logger.Info("HTTP server listening on", zap.String("port", cfg.HTTPPort))

		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			config.Logger.Fatal("HTTP server error", zap.Error(err))
		}
	}()

	// 10. Graceful shutdown
	<-ctx.Done()
	config.Logger.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		config.Logger.Error("graceful shutdown failed", zap.Error(err))
	}

	config.Logger.Info("server stopped")
}
