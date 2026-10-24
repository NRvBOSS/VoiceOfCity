package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NRvBOSS/citysesi/internal/config"
	"github.com/NRvBOSS/citysesi/internal/database"
	"github.com/NRvBOSS/citysesi/internal/db"
	"github.com/NRvBOSS/citysesi/internal/handler"
)

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")

		if origin == "http://localhost:3000" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Vary", "Origin")
			w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Konfiqurasiya yüklənmədi: %v", err)
	}

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)
	defer cancel()

	dbPool, err := database.Connect(ctx, cfg)
	if err != nil {
		log.Fatalf("Verilənlər bazasına qoşulmaq mümkün olmadı: %v", err)
	}
	defer dbPool.Close()

	queries := db.New(dbPool)
	issueHandler := handler.NewIssueHandler(queries)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)
	mux.HandleFunc("GET /issues", issueHandler.List)

	addr := ":" + cfg.ServerPort
	server := &http.Server{
		Addr:              addr,
		Handler:           corsMiddleware(mux),
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("PostgreSQL bağlantısı uğurludur")
	log.Printf("ŞəhərSəsi API %s portunda başladı", addr)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("API server xətası: %v", err)
	}
}
