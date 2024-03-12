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

func (p *ProductMapRepository) PartialUpdateProduct(index int, fields map[string]any) (err error) {
	for indexProduct, productItem := range (*p).Products {
		if productItem.ID == index {
			for field, value := range fields {
				switch field {
				case "name":
					name, ok := value.(string)
					if !ok {
						return internal.ErrorProductInvalidName
					}
					(*p).Products[indexProduct].Name = name
				case "quantity":
					quantity, ok := value.(int)
					if !ok {
						return internal.ErrorProductInvalidQuantity
					}
					(*p).Products[indexProduct].Quantity = quantity
				case "code":
					code, ok := value.(string)
					if !ok {
						return internal.ErrorProductInvalidCode
					}

					//verified if code is valid
					for _, productItem := range (*p).Products {
						if productItem.Code == code {
							return internal.ErrorProductDuplicated
						}
					}
					(*p).Products[indexProduct].Code = code
				case "published":
					published, ok := value.(bool)
					if !ok {
						return internal.ErrorProductInvalidPublished
					}
					(*p).Products[indexProduct].Published = published
				case "expiration":
					expiration, ok := value.(string)
					if !ok {
						return internal.ErrorProductInvalidExpiration
					}
					(*p).Products[indexProduct].Expiration = expiration
				case "price":
					price, ok := value.(float64)
					if !ok {
						return internal.ErrorProductInvalidPrice
					}
					(*p).Products[indexProduct].Price = price
				}
			}
			return nil
		}
	}
	return internal.ErrorProductNotFound
}

func (p *ProductMapRepository) DeleteProduct(id int) (err error) {
	for index, productItem := range (*p).Products {
		if productItem.ID == id {
			(*p).Products = append((*p).Products[:index], (*p).Products[index+1:]...)
			return nil
		}
	}
	return internal.ErrorProductNotFound
}
