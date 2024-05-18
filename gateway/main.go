package main

import (
	"log"
	"net/http"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common"
	"github.com/kaungmyathan22/golang-order-management-system-microservice/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr     = common.EnvString("HTTP_ADDR", ":8080")
	orderService = common.EnvString("ORDER_SERVICE", "localhost:3000")
)

func main() {
	conn, err := grpc.Dial(orderService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("failed to dial server %v", orderService)
	}
	defer conn.Close()
	log.Fatal("dialing order service at %v", orderService)
	c := api.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoute(mux)
	log.Printf("server is running at http://localhost%s\n", httpAddr)
	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatal("Failed to start http server.")
	}
}
