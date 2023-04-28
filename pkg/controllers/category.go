package controllers

import (
	"github.com/AzizRahimov/online_store_pc/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

type CategoryController struct {
	categoryService services.Category
}

func NewCategoryController(categoryService services.Category) CategoryController {
	return CategoryController{categoryService: categoryService}
}

func (c *CategoryController) GetAllCategory(e echo.Context) error {

	category, err := c.categoryService.GetAllCategory()
	if err != nil {
		return e.JSON(http.StatusBadRequest, err.Error())
	}

	return e.JSON(http.StatusOK, category)
}
