package services

import models "github.com/AzizRahimov/online_store_pc/model"

type Category interface {
	GetAllCategory() ([]models.Category, error)
}
