package service

import "github.com/lauramafoMeli/go-web/internal"

type ProductDefault struct {
	Repository internal.ProductsRepository
}

func NewProductDefault(repository internal.ProductsRepository) *ProductDefault {
	return &ProductDefault{
		Repository: repository,
	}
}

func (p *ProductDefault) GetAll() (products []internal.Product, err error) {
	return p.Repository.GetAll()
}

func (p *ProductDefault) GetProduct(id int) (product internal.Product, err error) {
	return p.Repository.GetProduct(id)
}

func (p *ProductDefault) GetProductsByPrice(price float64) (products []internal.Product, err error) {
	return p.Repository.GetProductsByPrice(price)
}

func (p *ProductDefault) SaveProduct(product *internal.Product) (err error) {
	return p.Repository.SaveProduct(product)
}

func (p *ProductDefault) UpdateProduct(product *internal.Product) (err error) {
	return p.Repository.UpdateProduct(product)
}

func (p *ProductDefault) PartialUpdateProduct(index int, fields map[string]any) (err error) {
	return p.Repository.PartialUpdateProduct(index, fields)
}
