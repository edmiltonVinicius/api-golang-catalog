package main

import (
	"context"
	"fmt"
	"log"
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
)

func main() {

	// 1. Load configuration
	cfg, err := config.LoadEnv()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// 2. Root context + OS signals
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	// 3. Postgres connection pool
	pool, err := pgxpool.New(ctx, cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("failed to create pgxpool: %v", err)
	}
	defer pool.Close()

	// 4. Basic health check
	if err := pool.Ping(ctx); err != nil {
		log.Fatalf("failed to ping postgres: %v", err)
	}
	log.Println("Postgres connection pool created successfully")

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
		log.Printf("HTTP server listening on :%s", cfg.HTTPPort)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error :%v", err)
		}
	}()

	// 10. Graceful shutdown
	<-ctx.Done()
	log.Println("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Printf("graceful shutdown failed: %v", err)
	}

	log.Println("server stopped")
}
