package services

import (
	models "github.com/AzizRahimov/online_store_pc/model"
	"github.com/AzizRahimov/online_store_pc/pkg/repository"
)

type CategoryServiceImp struct {
	repo repository.Category
}

func NewCategoryServiceImp(repo repository.Category) *CategoryServiceImp {
	return &CategoryServiceImp{repo: repo}
}

func (c *CategoryServiceImp) GetAllCategory() ([]models.Category, error) {
	parents, _ := c.repo.GetChildrenIDs(0)

	for i, p := range parents {

		ch1Arr, _ := c.repo.GetChildrenIDs(p.ID)
		parents[i].CategoryChildren = ch1Arr
		for i2, ch1 := range ch1Arr {
			ch2Arr, _ := c.repo.GetChildrenIDs(ch1.ID)
			ch1Arr[i2].CategoryChildren = ch2Arr
		}
	}

	return parents, nil
}
