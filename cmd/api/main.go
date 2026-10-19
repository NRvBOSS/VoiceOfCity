
package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/NRvBOSS/citysesi/internal/config"
	"github.com/NRvBOSS/citysesi/internal/database"
	"github.com/NRvBOSS/citysesi/internal/handler"
)

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

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)

	addr := ":" + cfg.ServerPort
	server := &http.Server{
		Addr:              addr,
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	log.Println("PostgreSQL bağlantısı uğurludur")
	log.Printf("ŞəhərSəsi API %s portunda başladı", addr)

	if err := server.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
		log.Fatalf("API server xətası: %v", err)
	}
}
