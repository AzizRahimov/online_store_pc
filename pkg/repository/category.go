package repository

import models "github.com/AzizRahimov/online_store_pc/model"

type Category interface {
	GetAllCategory() ([]models.Category, error)
	GetAllCategoryID() ([]int, error)
	GetParentIDs() ([]models.Category, error)
	GetChildrenIDs(id int) ([]models.Category, error)
}
