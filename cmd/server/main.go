package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go-app/internal/app"
)

func main() {
	appInstance := app.NewApp()
	http.HandleFunc("/", appInstance.HandleRequest)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := &http.Server{
		Addr:              ":" + port,
		ReadHeaderTimeout: 3 * time.Second,
	}

	fmt.Printf("Server starting on port %s...\n", port)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
