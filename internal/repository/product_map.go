package repository

import (
	"github.com/lauramafoMeli/go-web/internal"
)

type ProductMapRepository struct {
	Products []internal.Product `json:"products"`
	LastID   int                `json:"last_id"`
}

// GetAll implements internal.ProductsRepository.
func (p *ProductMapRepository) GetAll() (products []internal.Product, err error) {
	return p.Products, nil
}

func NewProductMapRepository(db []internal.Product, lastId int) *ProductMapRepository {
	return &ProductMapRepository{
		Products: db,
		LastID:   lastId,
	}
}

func (p *ProductMapRepository) GetProduct(id int) (product internal.Product, err error) {
	for _, product := range p.Products {
		if product.ID == id {
			return product, nil
		}
	}
	return internal.Product{}, internal.ErrorProductNotFound
}

func (p *ProductMapRepository) GetProductsByPrice(price float64) (products []internal.Product, err error) {
	var productsByPrice []internal.Product
	if price < 0 {
		return nil, internal.ErrorProductInvalidPrice
	}
	for _, product := range p.Products {
		if product.Price > price {
			productsByPrice = append(productsByPrice, product)
		}
	}
	return productsByPrice, nil
}

func (p *ProductMapRepository) SaveProduct(product *internal.Product) (err error) {
	for _, productItem := range (*p).Products {
		if productItem.Code == (*product).Code {
			return internal.ErrorProductDuplicated
		}
	}
	(*p).LastID++
	(*product).ID = (*p).LastID
	(*p).Products = append((*p).Products, *product)
	return nil
}

func (p *ProductMapRepository) UpdateProduct(product *internal.Product) (err error) {
	for _, productItem := range (*p).Products {
		if productItem.Code == (*product).Code {
			return internal.ErrorProductDuplicated
		}
	}
	for index, productItem := range (*p).Products {
		if productItem.ID == (*product).ID {
			(*p).Products[index] = *product
			return nil
		}
	}
	return internal.ErrorProductNotFound
}
