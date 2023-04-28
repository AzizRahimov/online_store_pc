package routes

import (
	"database/sql"
	"github.com/AzizRahimov/online_store_pc/pkg/controllers"
	"github.com/AzizRahimov/online_store_pc/pkg/repository"
	"github.com/AzizRahimov/online_store_pc/pkg/services"
	"github.com/labstack/echo/v4"
	"net/http"
)

func RunServerAndRoutes(db *sql.DB) {
	//server = gin.Default()
	//rc := server.Group("/api")
	e := echo.New()
	rc := e.Group("/api")
	categoryRep := repository.NewCategoryImp(db)
	categoryService := services.NewCategoryServiceImp(categoryRep)
	categoryController := controllers.NewCategoryController(categoryService)
	categoryControllerRoute := NewCategoryController(categoryController)
	categoryControllerRoute.CategoryRoute(rc)
	//e.GET("/", Home)
	e.Start(":9998")

}

func (r *CategoryController) Home(c echo.Context) error {
	return c.JSON(http.StatusOK, "Hello, World!")
}

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}
