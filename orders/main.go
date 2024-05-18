package main

import (
	"log"
	"net"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()
	l, err := net.Listen(grpcAddr, grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer l.Close()
	NewGRPCHandler(grpcServer)
	// store := NewStore()
	// svc := NewService(store)
	// svc.CreateOrder(context.Background())
	if err = grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
