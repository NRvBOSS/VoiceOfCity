package main

import (
	"log"
	"net/http"

	"github.com/NRvBOSS/citysesi/internal/config"
	"github.com/NRvBOSS/citysesi/internal/handler"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)

	addr := ":" + cfg.ServerPort
	log.Printf("ŞəhərSəsi API %s portunda başladı", addr)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
