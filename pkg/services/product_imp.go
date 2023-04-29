package services

import (
	models "github.com/AzizRahimov/online_store_pc/model"
	"github.com/AzizRahimov/online_store_pc/pkg/repository"
)

type ProductServiceImp struct {
	repo repository.ProductImp
}

func NewProductServiceImp(repo repository.ProductImp) *ProductServiceImp {
	return &ProductServiceImp{repo: repo}
}

func (p *ProductServiceImp) AddProduct(product models.Product) error {
	return p.repo.AddProduct(product)
}

func (p *ProductServiceImp) DeleteProduct(id int) error {
	return p.repo.DeleteProduct(id)

}

func (p *ProductServiceImp) UpdateProduct(product models.Product) error {
	return p.repo.UpdateProduct(product)

}

func (p *ProductServiceImp) GetAllProducts() ([]models.Product, error) {
	return p.repo.GetAllProducts()

}
func (p *ProductServiceImp) GetProduct(firstParam, secondParam string) ([]models.Product, error) {
	return p.repo.GetProduct(firstParam, secondParam)

}
