package routes

import (
	"github.com/AzizRahimov/online_store_pc/pkg/controllers"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	productRouteController controllers.ProductController
}

func NewProductController(productRouteController controllers.ProductController) *ProductController {
	return &ProductController{productRouteController: productRouteController}
}

func (r *ProductController) ProductRoute(group *echo.Group) {
	router := group.Group("/product")
	router.POST("/", r.productRouteController.AddProduct)
	router.DELETE("/delete/:id", r.productRouteController.DeleteProduct)
	router.PUT("/update/:id", r.productRouteController.UpdateProduct)
	router.GET("/all", r.productRouteController.GetAllProducts)
}
