package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/mfritschdotgo/techchallenge/internal/adapter/handler/dto"
	"github.com/mfritschdotgo/techchallenge/internal/core/service"
)

type OrderHandler struct {
	service *service.OrderService
}

func NewOrderHandler(s *service.OrderService) *OrderHandler {
	return &OrderHandler{
		service: s,
	}
}

// CreateOrder adds a new order to the store
// @Summary Add a new order
// @Description Adds a new order to the database with the given details.
// @Tags orders
// @Accept json
// @Produce json
// @Param		request	body		dto.CreateOrderRequest	true	"Order creation details"
// @Success 201 {object} domain.Order "Successfully created Order"
// @Failure 400 "Bad request if the Order data is invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /orders [post]
func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var orderDto dto.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&orderDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	order, err := h.service.CreateOrder(ctx, orderDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(order)
}

// GetOrderByID retrieves a order by its ID
// @Summary Get a order
// @Description Retrieves details of a order based on its unique ID.
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "order ID"
// @Success 200 {object} domain.Order "Successfully retrieved the order details"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 404 "Product not found if the ID does not match any order"
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrderByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid order ID", http.StatusBadRequest)
		return
	}

	order, err := h.service.GetOrderByID(ctx, id)
	if err != nil {
		http.Error(w, "Order not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(order)
}

// GetOrders retrieves a list of orders
// @Summary List orders
// @Description Retrieves a paginated list of orders
// @Tags orders
// @Accept json
// @Produce json
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Number of orders per page" default(10)
// @Success 200 {array} domain.Order "Successfully retrieved list of orders"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /orders [get]
func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	orders, err := h.service.GetOrders(ctx, page, size)
	if err != nil {
		http.Error(w, "Failed to retrieve orders", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(orders)
}
