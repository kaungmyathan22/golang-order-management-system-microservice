package main

import (
	"net/http"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common/api"
)

type handler struct {
	c api.OrderServiceClient
}

func NewHandler(c api.OrderServiceClient) *handler {
	return &handler{
		c: c,
	}
}

func (h *handler) registerRoute(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder)
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	h.c.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: r.URL.Query().Get("customerID"),
		// Items: ,
	})
}
