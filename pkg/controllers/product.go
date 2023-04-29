package controllers

import (
	models "github.com/AzizRahimov/online_store_pc/model"
	"github.com/AzizRahimov/online_store_pc/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Status struct {
	Message string `json:"message"`
}

type ProductController struct {
	productService services.Product
}

func NewProductController(productService services.Product) ProductController {
	return ProductController{productService: productService}
}
func (p *ProductController) AddProduct(c echo.Context) error {
	var product models.Product
	err := c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = p.productService.AddProduct(product)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, Status{Message: "продукт успешно добавлен"})

}

func (p *ProductController) DeleteProduct(c echo.Context) error {
	id := c.Param("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = p.productService.DeleteProduct(idInt)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, Status{Message: "Продукт успешно удален"})

}
func (p *ProductController) UpdateProduct(c echo.Context) error {
	var product models.Product
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	err = c.Bind(&product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())

	}
	product.ID = idInt
	err = p.productService.UpdateProduct(product)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, Status{Message: "Продукт успешно обновлен"})

}

func (p *ProductController) GetAllProducts(c echo.Context) error {
	products, err := p.productService.GetAllProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, products)

}