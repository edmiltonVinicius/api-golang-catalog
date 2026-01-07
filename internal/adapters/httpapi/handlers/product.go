package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/edmiltonVinicius/go-api-catalog/internal/adapters/httpapi/request"
	"github.com/edmiltonVinicius/go-api-catalog/internal/application/product"
	"github.com/edmiltonVinicius/go-api-catalog/internal/domain"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

type Handler struct {
	productService *product.Service
	validator      *validator.Validate
}

func NewHandler(productService *product.Service) *Handler {
	return &Handler{
		productService: productService,
		validator:      validator.New(),
	}
}

func (h *Handler) GetProduct(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	id := chi.URLParam(r, "id")

	if id == "" {
		writeError(w, http.StatusBadRequest, errors.New("id is required"))
		return
	}

	product, err := h.productService.GetProduct(ctx, id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"product": product,
	})
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var req request.CreateProduct
	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeError(w, http.StatusBadRequest, errors.New("Invalid json body"))
		return
	}

	if err := h.validator.Struct(req); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	domain, err := domain.NewProduct(domain.CreateProduct{
		Name:        req.Name,
		Price:       req.Price,
		Active:      *req.Active,
		Description: req.Description,
	})

	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	if err := h.productService.CreateProduct(ctx, *domain); err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	writeJSON(w, http.StatusOK, "Product sucessfully created")
}
