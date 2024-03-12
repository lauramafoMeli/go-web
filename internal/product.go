package internal

import "errors"

type Product struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Code       string  `json:"code_value"`
	Published  bool    `json:"is_published"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

type ProductsRepository interface {
	GetAll() (products []Product, err error)
	GetProduct(id int) (product Product, err error)
	GetProductsByPrice(price float64) (products []Product, err error)
	SaveProduct(product *Product) (err error)
}

type ProductService interface {
	GetAll() (products []Product, err error)
	GetProduct(id int) (product Product, err error)
	GetProductsByPrice(price float64) (products []Product, err error)
	SaveProduct(product *Product) (err error)
}

var (
	ErrorProductNotFound          = errors.New("product not found")
	ErrorProductDuplicated        = errors.New("product duplicated")
	ErrorProductInvalid           = errors.New("product invalid")
	ErrorProductInvalidID         = errors.New("product invalid id")
	ErrorProductInvalidName       = errors.New("product invalid name")
	ErrorProductInvalidQuantity   = errors.New("product invalid quantity")
	ErrorProductInvalidCode       = errors.New("product invalid code")
	ErrorProductInvalidPublished  = errors.New("product invalid published")
	ErrorProductInvalidExpiration = errors.New("product invalid expiration")
	ErrorProductInvalidPrice      = errors.New("product invalid price")
)
