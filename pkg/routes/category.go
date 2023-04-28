package routes

import (
	"github.com/AzizRahimov/online_store_pc/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type CategoryController struct {
	categoryRouteController controllers.CategoryController
}

func NewCategoryController(categoryRouteController controllers.CategoryController) *CategoryController {
	return &CategoryController{categoryRouteController: categoryRouteController}
}
func (r *CategoryController) CategoryRoute(group *echo.Group) {
	router := group.Group("/category")
	router.GET("/", r.categoryRouteController.GetAllCategory)
}
