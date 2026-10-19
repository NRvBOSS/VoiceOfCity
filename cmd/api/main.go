
package main

import (
	"log"
	"net/http"

	"github.com/NRvBOSS/citysesi/internal/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /health", handler.Health)

	log.Println("ŞəhərSəsi API :8080 portunda başladı")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}