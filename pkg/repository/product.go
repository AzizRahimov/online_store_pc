package repository

import models "github.com/AzizRahimov/online_store_pc/model"

type Product interface {
	AddProduct(product models.Product) error
	DeleteProduct(id int) error
	UpdateProduct(product models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProduct(firstParam, secondParam string) ([]models.Product, error)
}
