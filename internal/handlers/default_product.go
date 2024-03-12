package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/lauramafoMeli/go-web/internal"
	"github.com/lauramafoMeli/go-web/platform/tools"
)

type ProductRequest struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Code       string  `json:"code_value"`
	Published  bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

type DefaultProduct struct {
	Service internal.ProductService
}

func NewDefaultProduct(serviceProduct internal.ProductService) *DefaultProduct {
	return &DefaultProduct{
		Service: serviceProduct,
	}
}

// Get all products
func (p *DefaultProduct) GetAllProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		products, err := p.Service.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		json.NewEncoder(w).Encode(products)
	}
}

// Get product by id
func (p *DefaultProduct) GetProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		id := chi.URLParam(r, "id")
		idInt, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		product, err := p.Service.GetProduct(idInt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

// Get products by price
func (p *DefaultProduct) GetProductsByPrice() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		query := r.URL.Query()
		priceGt, err := strconv.ParseFloat(query.Get("priceGt"), 64)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		products, err := p.Service.GetProductsByPrice(priceGt)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(products)
	}
}

// Save product
func (p *DefaultProduct) SaveProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//validate product
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}

		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}

		err = tools.ValidateProductFields(bodyMap, "name", "quantity", "code_value", "expiration", "price")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := time.Parse("02/01/2006", fmt.Sprintf("%s", bodyMap["expiration"])); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid expiration format"))
			return
		}

		//if is published validate
		if _, ok := bodyMap["is_published"]; !ok || bodyMap["is_published"] == "" {
			bodyMap["is_published"] = false
		}

		//create product and add to products from bodyMap
		var product internal.Product
		aux, err := json.Marshal(bodyMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = json.Unmarshal(aux, &product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = (*p).Service.SaveProduct(&product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("201 - Created"))
	}
}

// Update product
func (p *DefaultProduct) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		//validate product
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}

		bodyMap := map[string]any{}
		if err := json.Unmarshal(bytes, &bodyMap); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}

		err = tools.ValidateProductFields(bodyMap, "name", "quantity", "code_value", "expiration", "price")
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}

		if _, err := time.Parse("02/01/2006", fmt.Sprintf("%s", bodyMap["expiration"])); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Invalid expiration format"))
			return
		}

		//if is published validate
		if _, ok := bodyMap["is_published"]; !ok || bodyMap["is_published"] == "" {
			bodyMap["is_published"] = false
		}

		//create product and add to products from bodyMap
		var product internal.Product
		aux, err := json.Marshal(bodyMap)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = json.Unmarshal(aux, &product)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		product.ID, err = strconv.Atoi(chi.URLParam(r, "id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("400 - Bad Request"))
			return
		}
		err = (*p).Service.UpdateProduct(&product)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("200 - OK"))
	}
}
