package main

import (
	"context"
	"log"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	api.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{
		api.UnimplementedOrderServiceServer{},
	}
	api.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(context.Context, *api.CreateOrderRequest) (*api.Order, error) {
	log.Println("New order recieved")
	o := api.Order{
		ID: "123",
	}
	return &o, nil
}
