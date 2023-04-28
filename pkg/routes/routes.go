package routes

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func RunServerAndRoutes() {
	//server = gin.Default()
	//rc := server.Group("/api")
	e := echo.New()
	e.GET("/", Home)
	e.Start(":9998")

}

func Home(c echo.Context) error {

	return c.String(http.StatusOK, "Hello, World!")
}
