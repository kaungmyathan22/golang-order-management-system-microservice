package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoute(mux)
	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal("Failed to start http server.")
	}
}
