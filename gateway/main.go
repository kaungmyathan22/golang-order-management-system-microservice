package main

import (
	"log"
	"net/http"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common"
)

var (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoute(mux)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server.")
	}
}
