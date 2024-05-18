package main

import (
	"net/http"

	"github.com/kaungmyathan22/golang-order-management-system-microservice/common"
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
	var items []*api.ItemsWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}
	h.c.CreateOrder(r.Context(), &api.CreateOrderRequest{
		CustomerID: r.URL.Query().Get("customerID"),
		Items:      items,
	})
}
