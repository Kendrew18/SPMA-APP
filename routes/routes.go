package routes

import (
	"Template-golang/controller/template_controller"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.DefaultCORSConfig))

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Project-NDL")
	})

	TMP := e.Group("/TMP")

	//NDL
	TMP.GET("/template", template_controller.Template_Controller)

	return e
}
