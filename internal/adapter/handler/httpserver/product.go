package httpserver

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"

	"github.com/mfritschdotgo/techchallenge/internal/adapter/handler/dto"
	"github.com/mfritschdotgo/techchallenge/internal/core/service"
)

type ProductHandler struct {
	service *service.ProductService
}

func NewProductHandler(s *service.ProductService) *ProductHandler {
	return &ProductHandler{
		service: s,
	}
}

// CreateProduct adds a new product to the store
// @Summary Add a new product
// @Description Adds a new product to the database with the given details.
// @Tags products
// @Accept json
// @Produce json
// @Param		request	body		dto.CreateProductRequest	true	"Product creation details"
// @Success 201 {object} domain.Product "Product successfully created"
// @Failure 400 "Bad request if the product data is invalid"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /products [post]
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var productDto dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.CreateProduct(ctx, productDto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

// UpdateProduct updates an existing product with the provided details.
// @Summary Update an existing product
// @Description Update product details in the database by ID.
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param		request	body		dto.CreateProductRequest	true	"Product object that needs to be updated"
// @Success 200 {object} domain.Product "Product successfully updated"
// @Failure 400 {string} string "Invalid input, Object is invalid"
// @Failure 404 {string} string "Product not found"
// @Failure 500 {string} string "Internal server error"
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	var productDto dto.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&productDto); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	product, err := h.service.UpdateProduct(ctx, id, productDto)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error updating product", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProductByID retrieves a product by its ID
// @Summary Get a product
// @Description Retrieves details of a product based on its unique ID.
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 {object} domain.Product "Successfully retrieved the product details"
// @Failure 400 "Bad request if the ID is not provided or invalid"
// @Failure 404 "Product not found if the ID does not match any product"
// @Router /products/{id} [get]
func (h *ProductHandler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	product, err := h.service.GetProductByID(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error retrieving product", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

// GetProducts retrieves a list of products
// @Summary List products
// @Description Retrieves a paginated list of products optionally filtered by category.
// @Tags products
// @Accept json
// @Produce json
// @Param category query string false "Filter products by category"
// @Param page query int false "Page number for pagination" default(1)
// @Param pageSize query int false "Number of products per page" default(10)
// @Success 200 {array} domain.Product "Successfully retrieved list of products"
// @Failure 500 "Internal server error if there is a problem on the server side"
// @Router /products [get]
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil || page < 1 {
		page = 1
	}

	size, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
	if err != nil || size <= 0 {
		size = 10
	}

	category := r.URL.Query().Get("category")

	products, err := h.service.GetProducts(ctx, category, page, size)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}

// DeleteProduct deletes a product by its ID
// @Summary Delete a product
// @Description Deletes a product based on its unique ID and returns a success message.
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Success 200 "Message indicating successful deletion"
// @Failure 400 "Bad request if the ID is not provided or is invalid"
// @Failure 404 "Product not found if the ID does not match any product"
// @Failure 500 "Internal server error if there is a problem deleting the product"
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")
	if id == "" {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	err := h.service.DeleteProduct(ctx, id)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			http.Error(w, "Product not found", http.StatusNotFound)
		} else {
			http.Error(w, "Error deleting product", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"message": "Product with ID " + id + " deleted successfully."}
	json.NewEncoder(w).Encode(response)
}
