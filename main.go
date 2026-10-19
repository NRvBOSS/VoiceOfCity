package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	response := HealthResponse{
		Status:  "ok",
		Service: "citysesi-api",
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("health response error: %v", err)
	}
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", healthHandler)

	log.Println("ŞəhərSəsi API :8080 portunda başladı")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
