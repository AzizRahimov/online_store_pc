package repository

import (
	"database/sql"
	"fmt"
	models "github.com/AzizRahimov/online_store_pc/model"
)

type CategoryImp struct {
	db *sql.DB
}

func NewCategoryImp(db *sql.DB) *CategoryImp {
	return &CategoryImp{db: db}
}

func (c *CategoryImp) GetAllCategory() ([]models.Category, error) {
	var categories []models.Category
	query := fmt.Sprintf("SELECT id, name, parent_id FROM categories ORDER BY id ASC")

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var category models.Category
		err := rows.Scan(&category.ID, &category.Title, &category.ParentID)
		if err != nil {
			return nil, err
		}

		categories = append(categories, category)

	}

	return categories, err

}

func (c *CategoryImp) GetAllCategoryID() ([]int, error) {
	var ids []int
	query := fmt.Sprintf("SELECT id FROM categories ORDER BY id ASC")

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}

		ids = append(ids, id)

	}

	return ids, err

}

func (c *CategoryImp) GetParentIDs() ([]models.Category, error) {
	var parentsCat []models.Category
	query := fmt.Sprintf("SELECT id, name, parent_id FROM categories WHERE parent_id = 0")

	rows, err := c.db.Query(query)
	if err != nil {
		return nil, err
	}

	for rows.Next() {

		var parentCat models.Category
		err := rows.Scan(&parentCat.ID, &parentCat.Title, &parentCat.ParentID)
		if err != nil {
			return nil, err
		}

		parentsCat = append(parentsCat, parentCat)

	}

	return parentsCat, err

}

func (c *CategoryImp) GetChildrenIDs(id int) ([]models.Category, error) {
	var categories []models.Category

	query := fmt.Sprintf("SELECT id, name, parent_id FROM categories  WHERE  parent_id = $1  ORDER BY id ASC")
	rows, err := c.db.Query(query, id)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var childrenID models.Category
		err := rows.Scan(&childrenID.ID, &childrenID.Title, &childrenID.ParentID)
		if err != nil {
			return nil, err
		}
		categories = append(categories, childrenID)

	}

	return categories, nil
}
