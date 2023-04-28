package services

import (
	"fmt"
	models "github.com/AzizRahimov/online_store_pc/model"
	"github.com/AzizRahimov/online_store_pc/pkg/repository"
)

type CategoryServiceImp struct {
	repo repository.Category
}

func NewCategoryServiceImp(repo repository.Category) *CategoryServiceImp {
	return &CategoryServiceImp{repo: repo}
}

func (c *CategoryServiceImp) GetAllCategory() ([]models.ChildrenID, error) {
	//ids, _ := c.repo.GetAllCategoryID()
	parentIds, _ := c.repo.GetParentIDs()

	categories, err := c.repo.GetAllCategory()
	if err != nil {
		return nil, err
	}
	var childrenIDs []models.ChildrenID
	var parentIDs []models.Category
	//1,2,3

	for _, parentID := range parentIds {
		fmt.Println("parentID", parentID)
		for _, category := range categories {
			if parentID.ID == category.ParentID {
				childrenIDs = append(childrenIDs, models.ChildrenID{
					ID:    category.ID,
					Title: category.Title,
				})

			}

		}

	}
	fmt.Println("parentIDs", parentIDs)

	return childrenIDs, err
}
